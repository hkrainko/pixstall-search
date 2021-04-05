//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	artistDeli "pixstall-search/app/artist/delivery/rabbitmq"
	artistUseCase "pixstall-search/app/artist/usecase"
	artworkDeli "pixstall-search/app/artwork/delivery/rabbitmq"
	artworkUseCase "pixstall-search/app/artwork/usecase"
	docRepo "pixstall-search/app/document/repo/elastic-search"
	openCommDeli "pixstall-search/app/open-commission/delivery/rabbitmq"
	openCommUseCase "pixstall-search/app/open-commission/usecase"
	searchDeli "pixstall-search/app/search/delivery/http"
	searchRepo "pixstall-search/app/search/repo/elastic-search"
	searchUseCase "pixstall-search/app/search/usecase"
)

func InitSearchController(host *docRepo.ElasticSearchHost) searchDeli.SearchController {
	wire.Build(
		searchDeli.NewSearchController,
		searchUseCase.NewSearchUseCase,
		searchRepo.NewElasticSearchSearchRepo,
	)
	return searchDeli.SearchController{}
}

func InitArtistMessageBroker(conn *amqp.Connection, host *docRepo.ElasticSearchHost) artistDeli.ArtistMessageBroker {
	wire.Build(
		artistDeli.NewRabbitMQArtistMessageBroker,
		artistUseCase.NewArtistUseCase,
		docRepo.NewElasticSearchDocumentRepo,
	)
	return artistDeli.ArtistMessageBroker{}
}

func InitArtworkMessageBroker(conn *amqp.Connection, host *docRepo.ElasticSearchHost) artworkDeli.ArtworkMessageBroker {
	wire.Build(
		artworkDeli.NewRabbitMQArtworkMessageBroker,
		artworkUseCase.NewArtworkUseCase,
		docRepo.NewElasticSearchDocumentRepo,
	)
	return artworkDeli.ArtworkMessageBroker{}
}

func InitOpenCommissionMessageBroker(conn *amqp.Connection, host *docRepo.ElasticSearchHost) openCommDeli.OpenCommissionMessageBroker {
	wire.Build(
		openCommDeli.NewRabbitMQOpenCommissionMessageBroker,
		openCommUseCase.NewOpenCommissionUseCase,
		docRepo.NewElasticSearchDocumentRepo,
		)
	return openCommDeli.OpenCommissionMessageBroker{}
}
