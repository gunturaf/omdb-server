package grpcservice

import (
	"log"
	"net"

	"github.com/gunturaf/omdb-server/infrastructure/grpcstub"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"google.golang.org/grpc"
)

func RunGRPCServer(port string, omdbService omdbservice.OMDBService) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	grpcstub.RegisterOmdbServer(grpcServer, NewGRPCService(omdbService))
	grpcServer.Serve(lis)
}
