package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"log"
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

	r := gin.Default()
	apiGroup := r.Group("/api")
	searchGroup := apiGroup.Group("/search")
	{
		searchGroup.GET("", )
	}

}
