package grpcservice_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/gunturaf/omdb-server/controllers/grpcservice"
	"github.com/gunturaf/omdb-server/entity"
	"github.com/gunturaf/omdb-server/infrastructure/repository/omdbservice"
)

var _ = Describe("Impl", func() {

	mockOMDBService := omdbservice.NewMock()

	Describe("Search", func() {
		Context("there's data", func() {
			It("ok", func() {
				mockOMDBService.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
					return &entity.OMDBSearchResult{
						Search: []entity.OMDBResultCompact{
							{
								IMDBID: "DavidBowie",
							},
						},
					}, nil
				}

				service := grpcservice.NewGRPCService(mockOMDBService)

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
			It("given False, error", func() {
				mockOMDBService.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
					return &entity.OMDBSearchResult{
						Response: "False",
					}, nil
				}

				service := grpcservice.NewGRPCService(mockOMDBService)

				_, err := service.Search(context.Background(), &entity.SearchRequest{
					Page:       1,
					Searchword: "Batman",
				})

				Expect(err).To(HaveOccurred())
			})

			It("error", func() {
				mockOMDBService.MockSearch = func(ctx context.Context, text string, page uint) (*entity.OMDBSearchResult, error) {
					return nil, nil
				}

				service := grpcservice.NewGRPCService(mockOMDBService)

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
				mockOMDBService.MockGetByID = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return &entity.OMDBResultSingle{
						Response: "True",
					}, nil
				}

				service := grpcservice.NewGRPCService(mockOMDBService)

				reply, err := service.Single(context.Background(), &entity.SingleRequest{
					Id: "DavidBowie",
				})

				Expect(err).NotTo(HaveOccurred())

				Expect(reply).NotTo(BeNil())
				Expect(reply.Response).To(Equal("True"))
			})
		})

		Context("no data", func() {
			It("given False, error", func() {
				mockOMDBService.MockGetByID = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return &entity.OMDBResultSingle{
						Response: "False",
					}, nil
				}

				service := grpcservice.NewGRPCService(mockOMDBService)

				_, err := service.Single(context.Background(), &entity.SingleRequest{
					Id: "DavidBowie",
				})

				Expect(err).To(HaveOccurred())
			})

			It("error", func() {
				mockOMDBService.MockGetByID = func(ctx context.Context, id string) (*entity.OMDBResultSingle, error) {
					return nil, nil
				}

				service := grpcservice.NewGRPCService(mockOMDBService)

				_, err := service.Single(context.Background(), &entity.SingleRequest{
					Id: "DavidBowie",
				})

				Expect(err).To(HaveOccurred())
			})
		})
	})

})
