package usecase

import (
	"context"

	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
)

type SingleUseCase interface {
	Single(ctx context.Context, id string) (*entity.OMDBResultSingle, error)
}

type SingleUseCaseImpl struct {
	omdbService omdbservice.OMDBService
}

func NewSingleUseCase(omdbService omdbservice.OMDBService) SingleUseCaseImpl {
	return SingleUseCaseImpl{
		omdbService: omdbService,
	}
}

func (s SingleUseCaseImpl) Single(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
	single, err := s.omdbService.GetByID(ctx, id)
	if err != nil || single == nil {
		return nil, ErrNoData
	}
	return single, nil
}

type MockSingleUseCase struct {
	MockSingle func(ctx context.Context, id string) (*entity.OMDBResultSingle, error)
}

func (m MockSingleUseCase) Single(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
	return m.MockSingle(ctx, id)
}

func NewMockSingleUseCase() MockSingleUseCase {
	return MockSingleUseCase{
		MockSingle: func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
			return nil, nil
		},
	}
}
