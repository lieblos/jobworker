package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/stoksc/jobworker/pkg/job"
	pb "github.com/stoksc/jobworker/pkg/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	m := job.NewManager()
	s := grpc.NewServer()

	pb.RegisterJobletServer(s, &server{jobService: &m})

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	go func() {
		<-sigint
		s.GracefulStop()
		// TODO if we want processes to not get clobbered when we stop,
		// we need to write a routine in the manager to release them all
		// call it here:
		// m.ReleaseAll()
	}()

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
