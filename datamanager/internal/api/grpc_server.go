package api

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func NewGrpcServer(
	address string, handler sensorpg.ReadingsServer, interceptors ...grpc.UnaryServerInterceptor,
) (*GrpcServer, error) {

	// Start server
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	// Create gRPC server
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors...),
	)

	// Register service handler to server
	sensorpg.RegisterReadingsServer(server, handler)
	//reflection.Register(server)

	//// Start listening to requests
	//// We put this in go routine in so that we don't block the main thread
	//// We block main thread with "<-ctx.Done()" so that we can have more control over
	//// closing open connections when program exits
	//go func() {
	//	if err := server.Serve(listener); err != nil {
	//		if !errors.Is(err, grpc.ErrServerStopped) {
	//			log.Printf("❌ gRPC server error: %v", err)
	//			stop()
	//		}
	//	}
	//}()

	return &GrpcServer{listener: listener, server: server}, nil
}

func (s *GrpcServer) Run(ctx context.Context) error {
	errChannel := make(chan error, 1)

	// server.Serve function is blocking operation. We don't want to block the main thread so we run it in go routine
	// In order to communicate between with other threads when error occurs in "server.Serve" thread we use go channel
	go func() {
		log.Printf("🚀 server listening at %v", s.listener.Addr())
		// Send error to Error Channel when it occurs
		errChannel <- s.server.Serve(s.listener)
	}()

	// This is similar to switch, if context gets stopped, we shut down the server
	// if error occurs in the above go routine, we handle than separately
	select {
	case <-ctx.Done():
		log.Println("Context canceled, shutting down server")
		s.server.GracefulStop()
		return nil
	case err := <-errChannel:
		if errors.Is(err, grpc.ErrServerStopped) {
			log.Println("Server stopped")
			return nil
		}

		return err
	}
}
