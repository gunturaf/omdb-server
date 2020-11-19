package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gunturaf/omdb-server/controllers/httpapi/handlers"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/gunturaf/omdb-server/usecase"
)

func RunServer(port string, omdbService omdbservice.OMDBService, searchUseCase usecase.SearchUseCase) {
	routes := mux.NewRouter()

	routes.Handle("/search", handlers.NewSearchHandler(searchUseCase))
	routes.Handle("/single/{id}", handlers.NewSingleHandler(omdbService))

	http.ListenAndServe("0.0.0.0:"+port, routes)
}
