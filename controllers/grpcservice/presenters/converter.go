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

func SingleToProto(inp *entity.OMDBResultSingle) *entity.SingleReply {
	out := entity.SingleReply{
		ImdbID:     inp.IMDBID,
		Actors:     inp.Actors,
		Awards:     inp.Awards,
		BoxOffice:  inp.BoxOffice,
		Country:    inp.Country,
		DVD:        inp.DVD,
		Director:   inp.Director,
		Genre:      inp.Genre,
		ImdbRating: inp.IMDBRating,
		ImdbVotes:  inp.IMDBVotes,
		Language:   inp.Language,
		Metascore:  inp.Metascore,
		Plot:       inp.Plot,
		Poster:     inp.Poster,
		Production: inp.Production,
		Rated:      inp.Rated,
		Ratings:    make([]*entity.SingleRating, 0),
		Released:   inp.Released,
		Response:   inp.Response,
		Runtime:    inp.Runtime,
		Title:      inp.Title,
		Type:       inp.Type,
		Website:    inp.Website,
		Writer:     inp.Writer,
		Year:       inp.Year,
	}

	for _, r := range inp.Ratings {
		out.Ratings = append(out.Ratings, &entity.SingleRating{
			Source: r.Source,
			Value:  r.Value,
		})
	}

	return &out
}
