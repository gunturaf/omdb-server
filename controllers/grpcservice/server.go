package grpcservice

import (
	"log"
	"net"

	"github.com/gunturaf/omdb-server/infrastructure/grpcstub"
	"github.com/gunturaf/omdb-server/usecase"
	"google.golang.org/grpc"
)

func RunGRPCServer(port string, searchUseCase usecase.SearchUseCase, singleUseCase usecase.SingleUseCase) {
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	grpcstub.RegisterOmdbServer(grpcServer, NewGRPCService(searchUseCase, singleUseCase))
	grpcServer.Serve(lis)
}
