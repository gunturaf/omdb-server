package presenters

import "github.com/gunturaf/omdb-server/entity"

func SearchResultToProto(inp *entity.OMDBSearchResult) *entity.SearchReply {
	out := entity.SearchReply{
		Search:       make([]*entity.SearchEntry, 0),
		Response:     inp.Response,
		TotalResults: inp.TotalResults,
	}

	for _, entry := range inp.Search {
		out.Search = append(out.Search, &entity.SearchEntry{
			ImdbID: entry.IMDBID,
			Poster: entry.Poster,
			Title:  entry.Title,
			Type:   entry.Type,
			Year:   entry.Year,
		})
	}

	return &out
}
