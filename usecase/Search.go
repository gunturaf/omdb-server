package usecase

import (
	"context"
	"errors"

	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository/mysqldb"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
)

var (
	ErrNoData = errors.New("no data")
)

type SearchUseCase interface {
	Search(ctx context.Context, searchWord string, page uint) (*entity.OMDBSearchResult, error)
}

type SearchUseCaseImpl struct {
	omdbService omdbservice.OMDBService
	mysqlRepo   mysqldb.MysqlDB
}

func NewSearchUseCase(omdbService omdbservice.OMDBService, mysqlRepo mysqldb.MysqlDB) SearchUseCaseImpl {
	return SearchUseCaseImpl{
		omdbService: omdbService,
		mysqlRepo:   mysqlRepo,
	}
}

func (s SearchUseCaseImpl) Search(ctx context.Context, searchWord string, page uint) (*entity.OMDBSearchResult, error) {
	go s.mysqlRepo.SaveSearchActivity(searchWord)

	response, err := s.omdbService.Search(ctx, searchWord, page)
	if err != nil || response == nil {
		return nil, ErrNoData
	}

	return response, nil
}

type MockSearchUseCase struct {
	MockSearch func(ctx context.Context, searchWord string, page uint) (*entity.OMDBSearchResult, error)
}

func (m MockSearchUseCase) Search(ctx context.Context, searchWord string, page uint) (*entity.OMDBSearchResult, error) {
	return m.MockSearch(ctx, searchWord, page)
}

func NewMockSearchUseCase() MockSearchUseCase {
	return MockSearchUseCase{
		MockSearch: func(ctx context.Context, searchWord string, page uint) (*entity.OMDBSearchResult, error) {
			return nil, nil
		},
	}
}
