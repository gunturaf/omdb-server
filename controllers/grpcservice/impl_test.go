package grpcservice_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gunturaf/omdb-server/controllers/grpcservice"
	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/usecase"
)

var _ = Describe("Impl", func() {

	mockSearchUseCase := usecase.NewMockSearchUseCase()
	mockSingleUseCase := usecase.NewMockSingleUseCase()

	Describe("Search", func() {
		Context("there's data", func() {
			It("ok", func() {
				mockSearchUseCase.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
					return &entity.OMDBSearchResult{
						Search: []entity.OMDBResultCompact{
							{
								IMDBID: "DavidBowie",
							},
						},
					}, nil
				}

				service := grpcservice.NewGRPCService(mockSearchUseCase, mockSingleUseCase)

				reply, err := service.Search(context.Background(), &entity.SearchRequest{
					Page:       1,
					Searchword: "Batman",
				})

				Expect(err).NotTo(HaveOccurred())

				Expect(reply).NotTo(BeNil())
				Expect(len(reply.GetSearch())).To(Equal(1))
			})
		})

		Context("no data", func() {
			It("error", func() {
				mockSearchUseCase.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
					return nil, nil
				}

				service := grpcservice.NewGRPCService(mockSearchUseCase, mockSingleUseCase)

				_, err := service.Search(context.Background(), &entity.SearchRequest{
					Page:       1,
					Searchword: "Batman",
				})

				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Single", func() {
		Context("there's data", func() {
			It("ok", func() {
				mockSingleUseCase.MockSingle = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return &entity.OMDBResultSingle{
						Response: "True",
					}, nil
				}

				service := grpcservice.NewGRPCService(mockSearchUseCase, mockSingleUseCase)

				reply, err := service.Single(context.Background(), &entity.SingleRequest{
					Id: "DavidBowie",
				})

				Expect(err).NotTo(HaveOccurred())

				Expect(reply).NotTo(BeNil())
				Expect(reply.Response).To(Equal("True"))
			})
		})

		Context("no data", func() {
			It("error", func() {
				mockSingleUseCase.MockSingle = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return nil, nil
				}

				service := grpcservice.NewGRPCService(mockSearchUseCase, mockSingleUseCase)

				_, err := service.Single(context.Background(), &entity.SingleRequest{
					Id: "DavidBowie",
				})

				Expect(err).To(HaveOccurred())
			})
		})
	})

})
