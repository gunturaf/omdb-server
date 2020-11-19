package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gunturaf/omdb-server/controllers/httpapi/presenters"
	"github.com/gunturaf/omdb-server/usecase"
)

type SingleHandler struct {
	singleUseCase usecase.SingleUseCase
}

func NewSingleHandler(singleUseCase usecase.SingleUseCase) SingleHandler {
	return SingleHandler{
		singleUseCase: singleUseCase,
	}
}

func (han SingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	single, err := han.singleUseCase.Single(r.Context(), id)
	if err != nil || single == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	presenters.WriteHTTPJSON(w, single)
}
