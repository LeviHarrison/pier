package image

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/archive"
)

// Build builds an image
func Build() {
	contents, err := archive.TarWithOptions(".", &archive.TarOptions{})
	if err != nil {
		fmt.Printf("error taring")
		os.Exit(1)
	}

	ctx := context.Background()
	build, err := client.ImageBuild(ctx, contents, types.ImageBuildOptions{})
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
