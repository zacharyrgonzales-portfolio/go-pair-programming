package cmd

import (
	"context"
	"fmt"

	"os"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
    "github.com/fatih/color"
)

func Dashboard() {
        // Define colors
        white := color.New(color.FgWhite).SprintFunc()
        brightBlue := color.New(color.FgHiBlue).SprintFunc()

        fmt.Println("Displaying the dashboard...\n")

        // Initialize a new API client and configure the Docker Env variables
        cli, err := client.NewClientWithOpts(client.FromEnv)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error Initializing docker client: %v\n", err)
            os.Exit(1)
        }
        defer cli.Close()

        // List Images
        images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error listing images from docker client: %v\n", err)
            os.Exit(1)
        }

        // Print Images with white text
        fmt.Println(white("Docker Images:"))
        for _, image := range images {
            fmt.Println(white(fmt.Sprintf("ID: %s, RepoTags: %s\n", image.ID, image.RepoTags)))
        }

        // List Containers
        containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error listing containers from docker client: %v\n", err)
            os.Exit(1)
        }
        
        // Print Containers
        fmt.Println(brightBlue("\nRunning Docker Containers:"))
        for _, container := range containers {
            fmt.Println(brightBlue(fmt.Sprintf("ID: %s, Image: %s, Names: %s, Status: %s\n", container.ID, container.Image, container.Names, container.Status)))
        }

}
