package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SingleHandler", func() {

	Describe("ServeHTTP", func() {
		mockOMDBService := omdbservice.NewMock()

		Context("there's any data", func() {
			It("ok", func() {
				id := "davidBowie"
				mockOMDBService.MockGetByID = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return &entity.OMDBResultSingle{
						OMDBResultCompact: entity.OMDBResultCompact{
							IMDBID: id,
						},
					}, nil
				}

				han := NewSingleHandler(mockOMDBService)

				w := httptest.NewRecorder()

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				r = mux.SetURLVars(r, map[string]string{
					"id": id,
				})

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			})
		})

		Context("no data", func() {
			It("error", func() {
				id := "davidBowie"
				mockOMDBService.MockGetByID = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return nil, nil
				}

				han := NewSingleHandler(mockOMDBService)

				w := httptest.NewRecorder()

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				r = mux.SetURLVars(r, map[string]string{
					"id": id,
				})

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})

})
