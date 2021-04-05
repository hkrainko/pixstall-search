// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/streadway/amqp"
	"pixstall-search/app/artist/delivery/rabbitmq"
	usecase2 "pixstall-search/app/artist/usecase"
	rabbitmq2 "pixstall-search/app/artwork/delivery/rabbitmq"
	usecase3 "pixstall-search/app/artwork/usecase"
	"pixstall-search/app/document/repo/elastic-search"
	rabbitmq3 "pixstall-search/app/open-commission/delivery/rabbitmq"
	usecase4 "pixstall-search/app/open-commission/usecase"
	"pixstall-search/app/search/delivery/http"
	elastic_search2 "pixstall-search/app/search/repo/elastic-search"
	"pixstall-search/app/search/usecase"
)

// Injectors from wire.go:

func InitSearchController(host *elastic_search.ElasticSearchHost) http.SearchController {
	repo := elastic_search2.NewElasticSearchSearchRepo(host)
	useCase := usecase.NewSearchUseCase(repo)
	searchController := http.NewSearchController(useCase)
	return searchController
}

func InitArtistMessageBroker(conn *amqp.Connection, host *elastic_search.ElasticSearchHost) rabbitmq.ArtistMessageBroker {
	repo := elastic_search.NewElasticSearchDocumentRepo(host)
	useCase := usecase2.NewArtistUseCase(repo)
	artistMessageBroker := rabbitmq.NewRabbitMQArtistMessageBroker(useCase, conn)
	return artistMessageBroker
}

func InitArtworkMessageBroker(conn *amqp.Connection, host *elastic_search.ElasticSearchHost) rabbitmq2.ArtworkMessageBroker {
	repo := elastic_search.NewElasticSearchDocumentRepo(host)
	useCase := usecase3.NewArtworkUseCase(repo)
	artworkMessageBroker := rabbitmq2.NewRabbitMQArtworkMessageBroker(useCase, conn)
	return artworkMessageBroker
}

func InitOpenCommissionMessageBroker(conn *amqp.Connection, host *elastic_search.ElasticSearchHost) rabbitmq3.OpenCommissionMessageBroker {
	repo := elastic_search.NewElasticSearchDocumentRepo(host)
	useCase := usecase4.NewOpenCommissionUseCase(repo)
	openCommissionMessageBroker := rabbitmq3.NewRabbitMQOpenCommissionMessageBroker(useCase, conn)
	return openCommissionMessageBroker
}
