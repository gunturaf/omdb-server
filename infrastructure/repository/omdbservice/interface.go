package omdbservice

import (
	"context"

	"github.com/gunturaf/omdb-server/entity"
)

type OMDBService interface {
	Search(ctx context.Context, text string, page uint) (entity.OMDBSearchResult, error)
	GetByID(ctx context.Context, id string) (entity.OMDBResultSingle, error)
}
