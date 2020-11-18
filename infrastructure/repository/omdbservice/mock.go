package omdbservice

import (
	"context"

	"github.com/gunturaf/omdb-server/entity"
)

type MockOMDB struct {
	MockSearch  func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error)
	MockGetByID func(ctx context.Context, id string) (*entity.OMDBResultSingle, error)
}

func NewMock() MockOMDB {
	return MockOMDB{
		MockSearch: func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
			return nil, nil
		},
		MockGetByID: func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
			return nil, nil
		},
	}
}

func (m MockOMDB) Search(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
	return m.MockSearch(ctx, text, page)
}
func (m MockOMDB) GetByID(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
	return m.MockGetByID(ctx, id)
}
