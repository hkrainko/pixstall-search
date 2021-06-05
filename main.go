package main

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	elastic_search "pixstall-search/app/document/repo/elastic-search"
)

func main() {
	// RabbitMQ
	rbMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer rbMQConn.Close()

	// Elasticsearch
	host := elastic_search.ElasticSearchHost{
		ApiPath:    "https://kai-dev-deployment.ent.ap-southeast-1.aws.found.io/api/as/v1/engines",
		SearchKey:  "search-gs7vdynu8fmqrpyz7vzhzf52",
		PrivateKey: "private-1xkbqfccp6xkrfo5t8j3vpia",
	}

	artistMsgBroker := InitArtistMessageBroker(rbMQConn, &host)
	go artistMsgBroker.StartArtistQueue()
	defer artistMsgBroker.StopArtistQueue()

	artworkMsgBroker := InitArtworkMessageBroker(rbMQConn, &host)
	go artworkMsgBroker.StartArtworkQueue()
	defer artworkMsgBroker.StopArtworkQueue()

	openCommMsgBroker := InitOpenCommissionMessageBroker(rbMQConn, &host)
	go openCommMsgBroker.StartOpenCommQueue()
	defer openCommMsgBroker.StopOpenCommQueue()

	r := gin.Default()
	apiGroup := r.Group("/api")
	searchGroup := apiGroup.Group("/search")
	{
		ctrl := InitSearchController(&host)
		searchGroup.GET("", ctrl.Search)
	}

	err = r.Run(":9006")
	print(err)
}
