package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	pb "github.com/stoksc/jobworker/pkg/proto"
)

type server struct {
	pb.UnimplementedJobletServer
}

// Create implements contract.JobletServer
func (s *server) Create(ctx context.Context, in *pb.Command) (*pb.Job, error) {
	return &pb.Job{Id: uuid.New().String(), Status: fmt.Sprintf("creating %v %v", in.Name, in.Args)}, nil
}

// Get implements contract.JobletServer
func (s *server) Get(ctx context.Context, in *pb.JobReq) (*pb.Job, error) {
	return &pb.Job{Id: in.GetId(), Status: fmt.Sprintf("retrieving %v", in.Id)}, nil
}

// Delete implements contract.JobletServer
func (s *server) Delete(ctx context.Context, in *pb.JobReq) (*pb.Job, error) {
	return &pb.Job{Id: in.GetId(), Status: fmt.Sprintf("stopping %v", in.Id)}, nil
}

// Tail implements contract.JobletServer
func (s *server) Tail(in *pb.JobReq, stream pb.Joblet_TailServer) error {
	for i := 0; i < 10; i++ {
		logLine := pb.Log{Line: fmt.Sprintf("logging %v", in.Id)}
		if err := stream.Send(&logLine); err != nil {
			return err
		}
	}
	return nil
}
