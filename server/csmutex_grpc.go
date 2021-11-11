package main

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/luks-itu/disys-mandatory-exercise-2/csmutex"
	"google.golang.org/grpc"
)

const (
	csDelay = 1000 * time.Millisecond
)

type CSMutexServer struct {
	csmutex.UnimplementedCSMutexServer
	// Data goes here
	mu          sync.Mutex
	currentNode int32
}

func newCSMutexServer() *CSMutexServer {
	return &CSMutexServer{
		currentNode: -1,
	}
}

func (s *CSMutexServer) RequestAccess(ctx context.Context, identifier *csmutex.Identifier) (*csmutex.Empty, error) {
	logI(fmt.Sprintf("Node %d requesting token", identifier.Id))
	s.mu.Lock()
	s.currentNode = identifier.Id
	return &csmutex.Empty{}, nil
}

func (s *CSMutexServer) ReleaseAccess(ctx context.Context, identifier *csmutex.Identifier) (*csmutex.Empty, error) {
	if identifier.Id != s.currentNode {
		err := fmt.Errorf("node %d denied release of lock due to current node not having token", identifier.Id)
		logE(err)
		return nil, err

	}
	s.mu.Unlock()
	s.currentNode = -1
	return &csmutex.Empty{}, nil
}

//cs method on server
func (s *CSMutexServer) PerformCriticalAction(ctx context.Context, action *csmutex.ActionDetails) (*csmutex.Empty, error) {
	if action.Id.Id != s.currentNode {
		err := fmt.Errorf("node %d denied access due to node not having token", action.Id.Id)
		logE(err)
		return nil, err
	}
	logI(fmt.Sprintf("Node %d entering critial section", action.Id.Id))
	time.Sleep(csDelay)
	logI(fmt.Sprintf("Node %d leaving critical section", action.Id.Id))
	return &csmutex.Empty{}, nil
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
