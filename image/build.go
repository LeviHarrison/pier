package image

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/archive"
	"github.com/leviharrison/pier"
)

// Build builds an image
func Build(target *pier.Target) {
	fmt.Println("building...")
	contents, err := archive.TarWithOptions(".", &archive.TarOptions{})
	if err != nil {
		fmt.Printf("Error TARing: %v", err)
		os.Exit(1)
	}

	ctx := context.Background()
	_, err = client.ImageBuild(ctx, contents, types.ImageBuildOptions{
		Dockerfile: target.Dockerfile,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
