package main

import (
	"context"
	"errors"
	"fmt"
	"io"

	pb "github.com/stoksc/jobworker/pkg/proto"
	"github.com/urfave/cli/v2"
)

func createJob(ctx context.Context, cl pb.JobletClient) func(*cli.Context) error {
	return func(c *cli.Context) error {
		name := c.Args().First()
		if name == "" {
			return errors.New("must provide a program to run")
		}

		cmd := &pb.Command{
			Name: name,
			Args: c.Args().Tail(),
		}

		job, err := cl.Create(ctx, cmd)
		if err != nil {
			return err
		}

		fmt.Println(job)
		return nil
	}
}

func getJob(ctx context.Context, cl pb.JobletClient) func(*cli.Context) error {
	return func(c *cli.Context) error {
		id := c.Args().First()
		if id == "" {
			return errors.New("must provide job id to retrieve")
		}

		jr := &pb.JobReq{
			Id: id,
		}

		job, err := cl.Get(ctx, jr)
		if err != nil {
			return err
		}

		fmt.Println(job)
		return nil
	}
}

func deleteJob(ctx context.Context, cl pb.JobletClient) func(*cli.Context) error {
	return func(c *cli.Context) error {
		id := c.Args().First()
		if id == "" {
			return errors.New("must provide job id to delete")
		}

		jr := &pb.JobReq{
			Id: id,
		}

		job, err := cl.Delete(ctx, jr)
		if err != nil {
			return err
		}

		fmt.Println(job)
		return nil
	}
}

func logs(ctx context.Context, cl pb.JobletClient) func(*cli.Context) error {
	return func(c *cli.Context) error {
		id := c.Args().First()
		if id == "" {
			return errors.New("must provide job id to get retrieve logs for")
		}

		jr := &pb.JobReq{
			Id: id,
		}

		tc, err := cl.Tail(ctx, jr)
		if err != nil {
			return err
		}

		for {
			resp, err := tc.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			fmt.Println(resp.Line)
		}

		return nil
	}
}
