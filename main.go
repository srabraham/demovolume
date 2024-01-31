package main

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"math/rand"
	"strconv"
)

// This program demonstrates that ryuk doesn't delete volumes that have nonempty names.
// To see this in action, you might want to watch the output of docker volume ls, e.g. by
//
//	while true; do docker volume ls; echo ""; sleep 1; done
//
// as you run this program, by
//
//	go run main.go
func main() {
	ctx := context.Background()

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "alpine:latest",
			Cmd:   []string{"echo", "yay"},
			Mounts: []testcontainers.ContainerMount{
				// ryuk will delete this volume
				testcontainers.VolumeMount("", "/mount1"),
				// ryuk won't delete this volume, since it has a name
				testcontainers.VolumeMount("this-wont-get-reaped-"+strconv.Itoa(rand.Int()), "/mount2"),
			},
			WaitingFor: wait.ForExit(),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	err = container.Terminate(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
