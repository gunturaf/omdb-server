package httpapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gunturaf/omdb-server/controllers/httpapi/handlers"
	"github.com/gunturaf/omdb-server/usecase"
)

func RunServer(port string, searchUseCase usecase.SearchUseCase, singleUseCase usecase.SingleUseCase) {
	routes := mux.NewRouter()

	routes.Handle("/search", handlers.NewSearchHandler(searchUseCase))
	routes.Handle("/single/{id}", handlers.NewSingleHandler(singleUseCase))

	http.ListenAndServe("0.0.0.0:"+port, routes)
}
