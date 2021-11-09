package main

import (
	"context"
	"fmt"
	"net"

	"github.com/luks-itu/disys-mandatory-exercise-2/csmutex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CSMutexServer struct {
	csmutex.UnimplementedCSMutexServer
	// Data goes here
}

func newCSMutexServer() *CSMutexServer {
	return &CSMutexServer{}
}

func (s *CSMutexServer) RequestAccess(ctx context.Context, identifier *csmutex.Identifier) (*csmutex.Empty, error) {
	//time.Sleep(3 * time.Second)
	logI(fmt.Sprintf("Node %d requesting token", identifier.Id))
	return &csmutex.Empty{}, nil
}

func (s *CSMutexServer) ReleaseAccess(ctx context.Context, empty *csmutex.Empty) (*csmutex.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseAccess not implemented")
}

func (s *CSMutexServer) PerformCriticalAction(ctx context.Context, action *csmutex.ActionDetails) (*csmutex.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformCriticalAction not implemented")
}

func startServer(address string) {
	logI(fmt.Sprintf("Starting server with address: %s", address))
	lis, err := net.Listen("tcp", address)
	if err != nil {
		logF(err.Error())
	}

	var opts []grpc.ServerOption
	// Server options here

	grpcServer := grpc.NewServer(opts...)

	csmutex.RegisterCSMutexServer(grpcServer, newCSMutexServer())

	logI("Server starting...")

	err = grpcServer.Serve(lis)
	if err != nil {
		logF(err.Error())
	}

}
