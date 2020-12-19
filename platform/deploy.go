package platform

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/digitalocean/godo"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
	"github.com/hashicorp/waypoint/builtin/docker"
)

// DeployConfig holds configuration for a deployment
type DeployConfig struct {
	Name             string `hcl:"name,optional"`
	Region           string `hcl:"region,optional"`
	InstanceSizeSlug string `hcl:"instance_size_slug,optional"`
	InstanceCount    int64  `hcl:"instance_count,optional"`

	HTTPPort int64  `hcl:"http_port,optional"`
	Path     string `hcl:"path,optional"`

	AccessToken string `hcl:"access_token,optional"`
}

// Platform is the Platform implementation for DigitalOcean
type Platform struct {
	config DeployConfig
	client *godo.Client
}

// Config implements Configurable
func (p *Platform) Config() (interface{}, error) {
	return &p.config, nil
}

// ConfigSet implement configurableNotify
func (p *Platform) ConfigSet(config interface{}) error {
	c, ok := config.(*DeployConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *DeployConfig as parameter")
	}

	tokenFromEnv := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if c.AccessToken == "" && tokenFromEnv != "" {
		c.AccessToken = tokenFromEnv
	}

	p.client = godo.NewFromToken(c.AccessToken)

	if c.Path == "" {
		c.Path = "/"
	}

	if c.HTTPPort == 0 {
		c.HTTPPort = 8080
	}

	return nil
}

// DeployFunc implements Builder
func (p *Platform) DeployFunc() interface{} {
	// return a function which will be called by Waypoint
	return p.deploy
}

// A BuildFunc does not have a strict signature, you can define the parameters
// you need based on the Available parameters that the Waypoint SDK provides.
// Waypoint will automatically inject parameters as specified
// in the signature at run time.
//
// Available input parameters:
// - context.Context
// - *component.Source
// - *component.JobInfo
// - *component.DeploymentConfig
// - *datadir.Project
// - *datadir.App
// - *datadir.Component
// - hclog.Logger
// - terminal.UI
// - *component.LabelSet

// In addition to default input parameters the registry.Artifact from the Build step
// can also be injected.
//
// The output parameters for BuildFunc must be a Struct which can
// be serialzied to Protocol Buffers binary format and an error.
// This Output Value will be made available for other functions
// as an input parameter.
// If an error is returned, Waypoint stops the execution flow and
// returns an error to the user.
func (p *Platform) deploy(
	ctx context.Context,
	ui terminal.UI,
	log hclog.Logger,
	src *component.Source,
	img *docker.Image) (*Deployment, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Deploying application")

	name := src.App
	if p.config.Name != "" {
		name = p.config.Name
	}

	appID, err := p.findExistingApp(name, u)
	if err != nil {
		return nil, err
	}

	registry, repository, regType := parseImage(img)

	spec := &godo.AppSpec{
		Name: name,
		Services: []*godo.AppServiceSpec{
			&godo.AppServiceSpec{
				Name:             name,
				InstanceSizeSlug: p.config.InstanceSizeSlug,
				InstanceCount:    p.config.InstanceCount,
				Image: &godo.ImageSourceSpec{
					RegistryType: regType,
					Registry:     registry,
					Repository:   repository,
					Tag:          img.Tag,
				},
				HTTPPort: p.config.HTTPPort,
				Routes: []*godo.AppRouteSpec{
					&godo.AppRouteSpec{
						Path: p.config.Path,
					},
				},
			},
		},
	}

	app := &godo.App{}
	if appID != "" {
		u.Update(fmt.Sprintf("Creating new deployment for existing application: %s (%s)", name, appID))
		appUpdateRequest := &godo.AppUpdateRequest{Spec: spec}
		app, _, err = p.client.Apps.Update(context.Background(), appID, appUpdateRequest)
		if err != nil {
			return nil, err
		}

	} else {
		u.Update(fmt.Sprintf("Creating new application: %s", name))
		appCreateRequest := &godo.AppCreateRequest{Spec: spec}
		app, _, err = p.client.Apps.Create(context.TODO(), appCreateRequest)
		if err != nil {
			return nil, err
		}
	}

	u.Update("Waiting for deployment to finish")
	err = p.waitForAppDeployment(app.ID, u)
	if err != nil {
		return nil, err
	}

	deployment := &Deployment{
		AppId:              app.ID,
		AppName:            name,
		DefaultIngress:     app.DefaultIngress,
		LiveUrl:            app.LiveURL,
		ActiveDeploymentId: app.ActiveDeployment.ID,
	}

	u.Step(terminal.StatusOK, fmt.Sprintf("Created App Platform deployment %s for %s", deployment.ActiveDeploymentId, name))
	ui.Output("\nDigitalOcean App Platform URL: %s", deployment.LiveUrl, terminal.WithSuccessStyle())

	return deployment, nil
}

func parseImage(img *docker.Image) (registry string, repository string, regType godo.ImageSourceSpecRegistryType) {
	repository = img.Image
	regType = godo.ImageSourceSpecRegistryType("UNSPECIFIED")

	imgParts := strings.Split(img.Image, "/")
	if len(imgParts) >= 3 && strings.Contains(imgParts[0], ".") {
		if imgParts[0] == "registry.digitalocean.com" {
			regType = godo.ImageSourceSpecRegistryType("DOCR")
			// The registry name must be left empty for the DOCR registry type.
			registry = ""
			// For DOCR we need to strip of the user name
			repository = strings.Join(imgParts[2:], "/")

			return registry, repository, regType
		}

		registry = imgParts[0]
		repository = strings.Join(imgParts[1:], "/")
	}

	return registry, repository, regType
}

func (p *Platform) findExistingApp(name string, u terminal.Status) (string, error) {
	list := []*godo.App{}
	opt := &godo.ListOptions{PerPage: 200, Page: 1}
	for {
		apps, resp, err := p.client.Apps.List(context.TODO(), opt)
		if err != nil {
			return "", err
		}

		list = append(list, apps...)

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return "", err
		}

		opt.Page = page + 1
	}

	var foundIDs []string
	for _, a := range list {
		if a.Spec.Name == name {
			foundIDs = append(foundIDs, a.ID)
		}
	}

	if len(foundIDs) >= 1 {
		u.Update(fmt.Sprintf("Found existing app for %s, ID: %s", name, foundIDs[0]))
		return foundIDs[0], nil
	}

	u.Update(fmt.Sprintf("No existing app found with name: %s", name))
	return "", nil
}

func (p *Platform) waitForAppDeployment(id string, u terminal.Status) error {
	tickerInterval := 10 //10s
	timeout := 1800      //1800s, 30min
	n := 0

	var deploymentID string
	ticker := time.NewTicker(time.Duration(tickerInterval) * time.Second)
	for range ticker.C {
		if n*tickerInterval > timeout {
			ticker.Stop()
			break
		}

		if deploymentID == "" {
			app, _, err := p.client.Apps.Get(context.Background(), id)
			if err != nil {
				return fmt.Errorf("Error trying to read app deployment state: %s", err)
			}

			if app.InProgressDeployment != nil {
				deploymentID = app.InProgressDeployment.ID
			}

		} else {
			deployment, _, err := p.client.Apps.GetDeployment(context.Background(), id, deploymentID)
			if err != nil {
				ticker.Stop()
				return fmt.Errorf("Error trying to read app deployment state: %s", err)
			}

			allSuccessful := deployment.Progress.SuccessSteps == deployment.Progress.TotalSteps
			if allSuccessful {
				ticker.Stop()
				return nil
			}

			if deployment.Progress.ErrorSteps > 0 {
				ticker.Stop()
				return fmt.Errorf("error deploying app (%s) (deployment ID: %s):\n%s", id, deployment.ID, godo.Stringify(deployment.Progress))
			}

			u.Update(fmt.Sprintf("Waiting for app (%s) deployment (%s) to become active. Phase: %s (%d/%d)",
				id, deployment.ID, deployment.Phase, deployment.Progress.SuccessSteps, deployment.Progress.TotalSteps))
		}

		n++
	}

	return fmt.Errorf("timeout waiting to app (%s) deployment", id)
}
