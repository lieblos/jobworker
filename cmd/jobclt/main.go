package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/stoksc/jobworker/pkg/proto"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cl := pb.NewJobletClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "creates a resource",
				Subcommands: []*cli.Command{
					{
						Name:   "job",
						Usage:  "creates a job resource",
						Action: createJob(ctx, cl),
					},
				},
			},
			{
				Name:  "get",
				Usage: "retrieves a resource",
				Subcommands: []*cli.Command{
					{
						Name:   "job",
						Usage:  "retrieves a job resource",
						Action: getJob(ctx, cl),
					},
				},
			},
			{
				Name:  "delete",
				Usage: "deletes a resource",
				Subcommands: []*cli.Command{
					{
						Name:   "job",
						Usage:  "deletes a job resource",
						Action: deleteJob(ctx, cl),
					},
				},
			},
			{
				Name:   "logs",
				Usage:  "retrieves the logs for a job",
				Action: logs(context.Background(), cl),
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("did not run: %v", err)
	}
}
