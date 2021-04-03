package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pixstall-search/app/artwork/delivery/rabbitmq/msg"
	"pixstall-search/domain/artwork"
	"time"
)

type ArtworkMessageBroker struct {
	useCase artwork.UseCase
	ch      *amqp.Channel
}

func NewRabbitMQArtworkMessageBroker(useCase artwork.UseCase, conn *amqp.Connection) ArtworkMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return ArtworkMessageBroker{
		useCase: useCase,
		ch:      ch,
	}
}

func (a ArtworkMessageBroker) StartArtistQueue() {
	//TODO
	q, err := a.ch.QueueDeclare(
		"artwork-event-to-search", // name
		true,                          // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = a.ch.QueueBind(
		q.Name,
		"artwork.event.#",
		"artist",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind queue %v", err)
	}

	msgs, err := a.ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)

			ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
			go func() {
				for {
					select {
					case <-ctx.Done():
						switch ctx.Err() {
						case context.DeadlineExceeded:
							log.Println("context.DeadlineExceeded")
						case context.Canceled:
							log.Println("context.Canceled")
						default:
							log.Println("default")
						}
						return // returning not to leak the goroutine
					}
				}
			}()

			switch d.RoutingKey {
			case "artwork.event.created":
				err := a.artworkCreated(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			case "artwork.event.updated":
				err := a.artworkUpdated(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			default:
				cancel()
			}


		}
	}()

	<-forever
}

func (a ArtworkMessageBroker) StopArtworkQueue() {
	err := a.ch.Close()
	if err != nil {
		log.Printf("StopArtworkQueue err %v", err)
	}
	log.Printf("StopArtworkQueue success")
}

func (a ArtworkMessageBroker) artworkCreated(ctx context.Context, body []byte) error {
	req := msg.CreatedArtwork{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	_, err = a.useCase.CreateArtwork(ctx, req.ToDomainArtworkCreator())
	if err != nil {
		return err
	}
	return nil
}

func (a ArtworkMessageBroker) artworkUpdated(ctx context.Context, body []byte) error {
	req := msg.UpdatedArtwork{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	_, err = a.useCase.UpdateArtwork(ctx, req.ToDomainArtworkUpdater())
	if err != nil {
		return err
	}
	return nil
}
