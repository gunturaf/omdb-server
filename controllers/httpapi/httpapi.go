package httpapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gunturaf/omdb-server/controllers/httpapi/handlers"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
)

func RunServer(port string, omdbService omdbservice.OMDBService) {
	routes := mux.NewRouter()

	routes.Handle("/search", handlers.NewSearchHandler(omdbService))

	go http.ListenAndServe("0.0.0.0:"+port, routes)

	fmt.Println("http api running at :" + port)
	select {}
}
