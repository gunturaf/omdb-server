package grpcservice

import (
	"log"
	"net"

	"github.com/gunturaf/omdb-server/infrastructure/grpcstub"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/gunturaf/omdb-server/usecase"
	"google.golang.org/grpc"
)

func RunGRPCServer(port string, omdbService omdbservice.OMDBService, searchUseCase usecase.SearchUseCase) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	grpcstub.RegisterOmdbServer(grpcServer, NewGRPCService(omdbService, searchUseCase))
	grpcServer.Serve(lis)
}
