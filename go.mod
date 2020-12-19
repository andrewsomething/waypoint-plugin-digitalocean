module github.com/andrewsomething/waypoint-plugin-digitalocean

go 1.14

require (
	github.com/digitalocean/godo v1.54.0
	github.com/golang/protobuf v1.4.3
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/waypoint v0.2.0
	github.com/hashicorp/waypoint-plugin-sdk v0.0.0-20201202203308-140d0145b90e
	google.golang.org/protobuf v1.25.0
)

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20200826173525-f9321e4c35a6
