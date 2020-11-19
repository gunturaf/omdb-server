package handlers

import (
	"net/http"
	"strconv"

	"github.com/gunturaf/omdb-server/controllers/httpapi/presenters"
	"github.com/gunturaf/omdb-server/usecase"
)

type SearchHandler struct {
	searchUseCase usecase.SearchUseCase
}

func NewSearchHandler(searchUseCase usecase.SearchUseCase) SearchHandler {
	return SearchHandler{
		searchUseCase: searchUseCase,
	}
}

func (han SearchHandler) getPageAndSearchWord(r *http.Request) (uint, string) {
	qs := r.URL.Query()

	page, err := strconv.Atoi(qs.Get("pagination"))
	if err != nil || page < 0 {
		page = 1
	}

	return uint(page), qs.Get("searchword")
}

func (han SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	page, searchword := han.getPageAndSearchWord(r)

	response, err := han.searchUseCase.Search(r.Context(), searchword, page)
	if err != nil || response == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	presenters.WriteHTTPJSON(w, response)
}
