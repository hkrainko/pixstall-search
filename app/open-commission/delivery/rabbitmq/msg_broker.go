package rabbitmq

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pixstall-search/app/open-commission/delivery/rabbitmq/msg"
	open_commission "pixstall-search/domain/open-commission"
	"time"
)

type OpenCommissionMessageBroker struct {
	useCase open_commission.UseCase
	ch      *amqp.Channel
}

func NewCommissionMessageBroker(useCase open_commission.UseCase, conn *amqp.Connection) OpenCommissionMessageBroker {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	err = ch.Qos(5, 0, false)
	if err != nil {
		log.Fatalf("Failed to set QoS %v", err)
	}

	return OpenCommissionMessageBroker{
		useCase: useCase,
		ch: ch,
	}
}

func (a OpenCommissionMessageBroker) StartArtistQueue() {
	//TODO
	q, err := a.ch.QueueDeclare(
		"open-comm-event-to-search", // name
		true,                        // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		nil,                         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}
	err = a.ch.QueueBind(
		q.Name,
		"open-comm.event.#",
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

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
			case "open-comm.event.created":
				err := a.openCommissionCreated(ctx, d.Body)
				if err != nil {
					//TODO: error handling, store it ?
				}
				cancel()
			case "open-comm.event.updated":
				err := a.openCommissionUpdated(ctx, d.Body)
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

func (a OpenCommissionMessageBroker) StopArtistQueue() {
	err := a.ch.Close()
	if err != nil {
		log.Printf("StopArtistQueue err %v", err)
	}
	log.Printf("StopArtistQueue success")
}

func (a OpenCommissionMessageBroker) openCommissionCreated(ctx context.Context, body []byte) error {
	req := msg.CreatedOpenCommission{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	_, err = a.useCase.CreateOpenCommission(ctx, req.ToDomainOpenCommissionCreator())
	if err != nil {
		return err
	}
	return nil
}

func (a OpenCommissionMessageBroker) openCommissionUpdated(ctx context.Context, body []byte) error {
	req := msg.UpdatedOpenCommission{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		return err
	}
	_, err = a.useCase.UpdateOpenCommission(ctx, req.ToDomainOpenCommissionUpdater())
	if err != nil {
		return err
	}
	return nil
}
