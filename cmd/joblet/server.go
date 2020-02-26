package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/stoksc/jobworker/pkg/job"
	pb "github.com/stoksc/jobworker/pkg/proto"
)

// JobService is the contract for a service that manages creating,
// deleting and monitoring Job resources.
type JobService interface {
	Create(name string, args ...string) (*job.Job, error)
	Get(id uuid.UUID) (*job.Job, error)
	Delete(id uuid.UUID) (string, error)
	Logs(id uuid.UUID) (chan (string), error)
}

type server struct {
	pb.UnimplementedJobletServer
	jobService JobService
}

// Create implements proto.JobletServer
func (s *server) Create(ctx context.Context, in *pb.Command) (*pb.Job, error) {
	j, err := s.jobService.Create(in.Name, in.Args...)
	if err != nil {
		return nil, err
	}

	return &pb.Job{Id: j.ID.String(), Status: <-j.Status}, nil
}

// Get implements proto.JobletServer
func (s *server) Get(ctx context.Context, in *pb.JobReq) (*pb.Job, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	j, err := s.jobService.Get(id)
	if err != nil {
		return nil, err
	}

	return &pb.Job{Id: j.ID.String(), Status: <-j.Status}, nil
}

// Delete implements proto.JobletServer
func (s *server) Delete(ctx context.Context, in *pb.JobReq) (*pb.Job, error) {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, err
	}

	stat, err := s.jobService.Delete(id)
	if err != nil {
		return nil, err
	}

	return &pb.Job{Id: id.String(), Status: stat}, nil
}

// Tail implements proto.JobletServer
func (s *server) Tail(in *pb.JobReq, stream pb.Joblet_TailServer) error {
	id, err := uuid.Parse(in.Id)
	if err != nil {
		return err
	}

	logs, err := s.jobService.Logs(id)
	if err != nil {
		return err
	}

	for l := range logs {
		if err := stream.Send(&pb.Log{Line: l}); err != nil {
			return err
		}
	}
	return nil
}
