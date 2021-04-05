package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
	elastic_search "pixstall-search/app/document/repo/elastic-search"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// RabbitMQ
	rbMQConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer rbMQConn.Close()
	ch, err := rbMQConn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel %v", err)
	}

	host := elastic_search.ElasticSearchHost{
		ApiPath:    "http://localhost:3002/api/as/v1/engines",
		SearchKey:  "search-21ovx6yqaffz92sxw8qpu23p",
		PrivateKey: "private-jxcq12x5tuko6rkbh28gkbdk",
	}

	r := gin.Default()
	apiGroup := r.Group("/api")
	searchGroup := apiGroup.Group("/search")
	{
		ctrl := InitSearchController(&host)
		searchGroup.GET("", ctrl.Search)
	}

}
