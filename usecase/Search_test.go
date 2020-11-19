package usecase_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository/mysqldb"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/gunturaf/omdb-server/usecase"
)

var _ = Describe("Search", func() {

	mockOMDBService := omdbservice.NewMock()
	mockMysqlDB := mysqldb.NewMock()

	Context("there's data", func() {
		It("ok", func() {
			mockOMDBService.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
				return &entity.OMDBSearchResult{}, nil
			}
			use := usecase.NewSearchUseCase(mockOMDBService, mockMysqlDB)

			resp, err := use.Search(context.Background(), "Bowie", 1)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).NotTo(BeNil())
		})
	})

	Context("no data", func() {
		It("error", func() {
			mockOMDBService.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
				return nil, nil
			}
			use := usecase.NewSearchUseCase(mockOMDBService, mockMysqlDB)

			resp, err := use.Search(context.Background(), "Bowie", 1)
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})

})
