# DigitalOcean Waypoint Plugin

This is a prototype DigitalOcean Waypoint plugin. It currently has _very_ basic
support for deploying a Docker image from a DigitalOcean Container Registry to
DigitalOcean App Platform. **It should be considered experimental.**

## Configuration

A Waypoint configuration of this might look like:

```hcl
project = "example-nodejs"

app "example-nodejs" {
  build {
    use "pack" {}
    registry {
        use "docker" {
          image = "registry.digitalocean.com/<username>/example-nodejs"
          tag   = "latest"
        }
    }
 }

  deploy {
    use "digitalocean" {
    }
  }
}
```

The following configuration options are supported. They are all optional.

* access_token - Required if `DIGITALOCEAN_ACCESS_TOKEN` is not set
* name - Defaults to the app's name
* region - Defaults to nearest region
* instance_size_slug - Defaults to `basic-xxs`
* instance_count - Default to `1`
* http_port - Default to `8080`
* path - Default to `/`


## Development

### Building

To build the plugin, run:

```shell
make
```

This will regenerate the protos and build binaries for multiple platforms.

```shell
Build Protos
protoc -I . --go_out=plugins=grpc:. --go_opt=paths=source_relative ./platform/output.proto

Compile Plugin
# Clear the output
rm -rf ./bin
GOOS=linux GOARCH=amd64 go build -o ./bin/linux_amd64/waypoint-plugin-digitalocean ./main.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin_amd64/waypoint-plugin-digitalocean ./main.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows_amd64/waypoint-plugin-digitalocean.exe ./main.go
GOOS=windows GOARCH=386 go build -o ./bin/windows_386/waypoint-plugin-digitalocean.exe ./main.go
```

### Building with Docker

To build plugins for release you can use the `build-docker` Makefile target, this will
build your plugin for all architectures and create zipped artifacts which can be uploaded
to an artifact manager such as GitHub releases.

The built artifacts will be output in the `./releases` folder.

```shell
make build-docker

rm -rf ./releases
DOCKER_BUILDKIT=1 docker build --output releases --progress=plain .
#1 [internal] load .dockerignore
#1 transferring context: 2B done
#1 DONE 0.0s

#...

#14 [export_stage 1/1] COPY --from=build /go/plugin/bin/*.zip .
#14 DONE 0.1s

#15 exporting to client
#15 copying files 36.45MB 0.1s done
#15 DONE 0.1s
```

### Building and releasing with GitHub Actions

When cloning the template a default GitHub Action is created at the path `.github/workflows/build-plugin.yaml`. You can use this action to automatically build and release your plugin.

The action has two main phases:
1. **Build** - This phase builds the plugin binaries for all the supported architectures. It is triggered when pushing
   to a branch or on pull requests.
1. **Release** - This phase creates a new GitHub release containing the built plugin. It is triggered when pushing tags
   which starting with `v`, for example `v0.1.0`.

You can enable this action by clicking on the `Actions` tab in your GitHub repository and enabling GitHub Actions.