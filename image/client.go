package image

import (
	"fmt"
	"os"

	docker "github.com/docker/docker/client"
)

var client *docker.Client

// NewClient initializes a new Docker client
func NewClient() {
	c, err := docker.NewEnvClient()
	if err != nil {
		fmt.Printf("Error initializing Docker client: %v\n", err)
		os.Exit(1)
	}

	client = c
}
