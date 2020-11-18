package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gunturaf/omdb-server/controllers/httpapi/presenters"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
)

type SingleHandler struct {
	omdbService omdbservice.OMDBService
}

func NewSingleHandler(omdbService omdbservice.OMDBService) SingleHandler {
	return SingleHandler{
		omdbService: omdbService,
	}
}

func (han SingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	single, err := han.omdbService.GetByID(r.Context(), id)
	if err != nil || single == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	presenters.WriteHTTPJSON(w, single)
}
