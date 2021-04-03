package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pixstall-search/app/artist/delivery/rabbitmq/msg"
	"pixstall-search/domain/artist"
	"time"
)

type ArtistMessageBroker struct {
	useCase artist.UseCase
	ch      *amqp.Channel
}

func NewRabbitMQArtistMessageBroker(useCase artist.UseCase, conn *amqp.Connection) ArtistMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return ArtistMessageBroker{
		useCase: useCase,
		ch:      ch,
	}
}

func (a ArtistMessageBroker) StartArtistQueue() {
	//TODO
	q, err := a.ch.QueueDeclare(
		"artist-event-to-search", // name
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
		"artist.event.#",
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
			case "artist.event.created":
				err := a.artistCreated(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			case "artist.event.updated":
				err := a.artistUpdated(ctx, d.Body)
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

func (a ArtistMessageBroker) StopArtistQueue() {
	err := a.ch.Close()
	if err != nil {
		log.Printf("StopArtistQueue err %v", err)
	}
	log.Printf("StopArtistQueue success")
}

func (a ArtistMessageBroker) artistCreated(ctx context.Context, body []byte) error {
	req := msg.CreatedArtist{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	_, err = a.useCase.CreateArtist(ctx, req.ToDomainArtistCreator())
	if err != nil {
		return err
	}
	return nil
}

func (a ArtistMessageBroker) artistUpdated(ctx context.Context, body []byte) error {
	req := msg.UpdatedArtist{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	_, err = a.useCase.UpdateArtist(ctx, req.ToDomainArtistUpdater())
	if err != nil {
		return err
	}
	return nil
}
