package usecase_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
	"github.com/gunturaf/omdb-server/usecase"
)

var _ = Describe("Single", func() {
	mockOMDBService := omdbservice.NewMock()

	Context("there's data", func() {
		It("ok", func() {
			mockOMDBService.MockGetByID = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
				return &entity.OMDBResultSingle{}, nil
			}
			use := usecase.NewSingleUseCase(mockOMDBService)

			resp, err := use.Single(context.Background(), "Bowie")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp).NotTo(BeNil())
		})
	})

	Context("no data", func() {
		It("error", func() {
			mockOMDBService.MockGetByID = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
				return nil, nil
			}
			use := usecase.NewSingleUseCase(mockOMDBService)

			resp, err := use.Single(context.Background(), "Bowie")
			Expect(err).To(HaveOccurred())
			Expect(resp).To(BeNil())
		})
	})
})
