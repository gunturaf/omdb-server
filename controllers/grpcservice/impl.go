package grpcservice

import (
	"context"

	"github.com/gunturaf/omdb-server/controllers/grpcservice/presenters"
	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/grpcstub"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServiceImpl struct {
	grpcstub.UnimplementedOmdbServer

	omdbService omdbservice.OMDBService
}

func NewGRPCService(omdbService omdbservice.OMDBService) GRPCServiceImpl {
	return GRPCServiceImpl{
		omdbService: omdbService,
	}
}

func (impl GRPCServiceImpl) Search(ctx context.Context, r *entity.SearchRequest) (*entity.SearchReply, error) {
	response, err := impl.omdbService.Search(ctx, r.GetSearchword(), uint(r.GetPage()))
	if err != nil || response == nil {
		return nil, status.Errorf(codes.NotFound, "not found")
	}

	return presenters.SearchResultToProto(response), nil
}

func (impl GRPCServiceImpl) Single(ctx context.Context, r *entity.SingleRequest) (*entity.SingleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Single not implemented")
}
