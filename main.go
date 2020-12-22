package main

import (
	"github.com/andrewsomething/waypoint-plugin-digitalocean/platform"
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

func main() {
	sdk.Main(sdk.WithComponents(
		&platform.Platform{},
		&platform.Releaser{},
	))
}
