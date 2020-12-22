package platform

import (
	"context"
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/waypoint-plugin-sdk/component"
	"github.com/hashicorp/waypoint-plugin-sdk/terminal"
)

// ReleaseConfig is the configuration structure for the Releaser.
type ReleaseConfig struct {
	AccessToken string `hcl:"access_token,optional"`
}

// Releaser is the ReleaseManager implementation
type Releaser struct {
	config ReleaseConfig
	client *godo.Client
}

// Config implements Configurable
func (r *Releaser) Config() (interface{}, error) {
	return &r.config, nil
}

// ConfigSet implements ConfigurableNotify
func (r *Releaser) ConfigSet(config interface{}) error {
	c, ok := config.(*ReleaseConfig)
	if !ok {
		// The Waypoint SDK should ensure this never gets hit
		return fmt.Errorf("Expected *ReleaseConfig as parameter")
	}

	tokenFromEnv := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if c.AccessToken == "" && tokenFromEnv != "" {
		c.AccessToken = tokenFromEnv
	}

	r.client = godo.NewFromToken(c.AccessToken)

	return nil
}

// ReleaseFunc implements component.ReleaseManager
func (r *Releaser) ReleaseFunc() interface{} {
	// return a function which will be called by Waypoint
	return r.release
}

// DestroyFunc implements component.Destroyer
func (r *Releaser) DestroyFunc() interface{} {
	return r.destroy
}

func (r *Release) URL() string { return r.Url }

func (r *Releaser) release(ctx context.Context, ui terminal.UI, deployment *Deployment, log hclog.Logger) (*Release, error) {
	u := ui.Status()
	defer u.Close()
	u.Update("Release application")

	log.Debug(fmt.Sprintf("Registering release for app: %s", deployment.AppId))
	release := &Release{
		AppId: deployment.AppId,
		Url:   deployment.Url,
	}

	return release, nil
}

func (r *Releaser) destroy(ctx context.Context, ui terminal.UI, release *Release, log hclog.Logger) error {
	u := ui.Status()
	defer u.Close()
	u.Update("Destroying application")

	log.Debug(fmt.Sprintf("Destroying application: %s !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!", release.AppId))
	_, err := r.client.Apps.Delete(context.TODO(), release.AppId)

	return err
}

var _ component.ReleaseManager = (*Releaser)(nil)
var _ component.Destroyer = (*Releaser)(nil)
