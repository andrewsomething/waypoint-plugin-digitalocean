package main

import (
	"github.com/andrewsomething/waypoint-plugin-digitalocean/platform"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	sdk.Main(sdk.WithComponents(
		// Comment out any components which are not
		// required for your plugin
		// &builder.Builder{},
		// &registry.Registry{},
		&platform.Platform{},
		// &release.ReleaseManager{},
	))
}
