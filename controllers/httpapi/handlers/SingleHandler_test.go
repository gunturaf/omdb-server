package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SingleHandler", func() {

	Describe("ServeHTTP", func() {
		mockSingleUseCase := usecase.NewMockSingleUseCase()

		Context("there's any data", func() {
			It("ok", func() {
				id := "davidBowie"
				mockSingleUseCase.MockSingle = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return &entity.OMDBResultSingle{
						OMDBResultCompact: entity.OMDBResultCompact{
							IMDBID: id,
						},
					}, nil
				}

				han := NewSingleHandler(mockSingleUseCase)

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
				mockSingleUseCase.MockSingle = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return nil, nil
				}

				han := NewSingleHandler(mockSingleUseCase)

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
