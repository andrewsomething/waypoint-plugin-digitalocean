package platform

import (
	"testing"

	"github.com/digitalocean/godo"
	"github.com/hashicorp/waypoint/builtin/docker"
)

func TestParseImage(t *testing.T) {
	tests := []struct {
		image   *docker.Image
		reg     string
		repo    string
		regType godo.ImageSourceSpecRegistryType
	}{
		{
			&docker.Image{Image: "registry.digitalocean.com/foo/bar"},
			"", "bar", godo.ImageSourceSpecRegistryType("DOCR"),
		},
		{
			&docker.Image{Image: "registry.digitalocean.com/foo/bar/baz"},
			"", "bar/baz", godo.ImageSourceSpecRegistryType("DOCR"),
		},
		{
			&docker.Image{Image: "example.com/foo/bar"},
			"example.com", "foo/bar", godo.ImageSourceSpecRegistryType("UNSPECIFIED"),
		},
		{
			&docker.Image{Image: "foo/bar"},
			"", "foo/bar", godo.ImageSourceSpecRegistryType("UNSPECIFIED")},
		{
			&docker.Image{Image: "foo/bar/baz"},
			"", "foo/bar/baz", godo.ImageSourceSpecRegistryType("UNSPECIFIED")},
	}

	for _, tt := range tests {
		t.Run(tt.image.Image, func(t *testing.T) {
			registry, repository, regType := parseImage(tt.image)
			if registry != tt.reg {
				t.Errorf("got %q, want %q", registry, tt.reg)
			}
			if repository != tt.repo {
				t.Errorf("got %q, want %q", repository, tt.repo)
			}
			if regType != tt.regType {
				t.Errorf("got %q, want %q", regType, tt.regType)
			}
		})
	}
}
