package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/gunturaf/omdb-server/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SearchHandler", func() {

	Describe("getPageAndSearchWord", func() {

		mockSearchUseCase := usecase.NewMockSearchUseCase()

		Context("given pagination and searchword", func() {
			It("yield correct", func() {
				han := NewSearchHandler(mockSearchUseCase)

				r, err := http.NewRequest(http.MethodGet, "/?pagination=2&searchword=Batman", nil)
				Expect(err).NotTo(HaveOccurred())

				page, searchWord := han.getPageAndSearchWord(r)
				Expect(page).To(Equal(uint(2)))
				Expect(searchWord).To(Equal("Batman"))
			})
		})
		Context("given pagination only", func() {
			It("yield correct", func() {
				han := NewSearchHandler(mockSearchUseCase)

				r, err := http.NewRequest(http.MethodGet, "/?pagination=2", nil)
				Expect(err).NotTo(HaveOccurred())

				page, searchWord := han.getPageAndSearchWord(r)
				Expect(page).To(Equal(uint(2)))
				Expect(searchWord).To(Equal(""))
			})
		})
		Context("given searchword only", func() {
			It("yield correct", func() {
				han := NewSearchHandler(mockSearchUseCase)

				r, err := http.NewRequest(http.MethodGet, "/?searchword=Batman", nil)
				Expect(err).NotTo(HaveOccurred())

				page, searchWord := han.getPageAndSearchWord(r)
				Expect(page).To(Equal(uint(1)))
				Expect(searchWord).To(Equal("Batman"))
			})
		})
	})

	Describe("ServeHTTP", func() {
		Context("there's data", func() {
			It("return http.StatusOK", func() {
				mockSearchUseCase := usecase.NewMockSearchUseCase()

				mockSearchUseCase.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
					return &entity.OMDBSearchResult{}, nil
				}

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				w := httptest.NewRecorder()

				han := NewSearchHandler(mockSearchUseCase)

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
			})
		})
		Context("no data", func() {
			It("return http.StatusNotFound", func() {
				mockSearchUseCase := omdbservice.NewMock()

				mockSearchUseCase.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
					return nil, nil
				}

				r, err := http.NewRequest(http.MethodGet, "/", nil)
				Expect(err).NotTo(HaveOccurred())

				w := httptest.NewRecorder()

				han := NewSearchHandler(mockSearchUseCase)

				han.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusNotFound))
			})
		})

	})

})
