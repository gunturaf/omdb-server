package handlers

import (
	"net/http"
	"strconv"

	"github.com/gunturaf/omdb-server/controllers/httpapi/presenters"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
)

type SearchHandler struct {
	omdbService omdbservice.OMDBService
}

func NewSearchHandler(omdbService omdbservice.OMDBService) SearchHandler {
	return SearchHandler{
		omdbService: omdbService,
	}
}

func (han SearchHandler) getPageAndSearchWord(r *http.Request) (uint, string) {
	qs := r.URL.Query()

	page, err := strconv.Atoi(qs.Get("pagination"))
	if err != nil || page < 0 {
		page = 1
	}
	searchword := qs.Get("searchword")

	return uint(page), searchword
}

func (han SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page, searchword := han.getPageAndSearchWord(r)

	response, err := han.omdbService.Search(r.Context(), searchword, page)
	if err != nil || response == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	presenters.WriteHTTPJSON(w, response)
}
