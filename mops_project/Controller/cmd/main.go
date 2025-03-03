// controller/main.go
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DataMessage описывает структуру отправляемого пакета
type DataMessage struct {
	DeviceID  int     `json:"device_id"`
	FieldA    float64 `json:"field_a"`
	FieldB    float64 `json:"field_b"`
	RuleField int     `json:"rule_fld"`
}

func main() {
	var amqpURL, queueName string
	var mongoURI string
	var mongoDBName, mongoCollName string
	var httpServerAddr string

	flag.StringVar(&amqpURL, "amqp", "amqp://guest:guest@localhost:5672/", "AMQP URL")
	flag.StringVar(&queueName, "queue", "iot_data", "RabbitMQ queue name")
	flag.StringVar(&mongoURI, "mongo-uri", "mongodb://localhost:27017", "MongoDB URI")
	flag.StringVar(&mongoDBName, "mongo-db", "iot_db", "MongoDB database name")
	flag.StringVar(&mongoCollName, "mongo-coll", "messages", "MongoDB collection name")
	flag.StringVar(&httpServerAddr, "server", ":8081", "HTTP server address")
	flag.Parse()

	// Подключаемся к MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
	}()
	collection := client.Database(mongoDBName).Collection(mongoCollName)

	// Подключаемся к RabbitMQ
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Убеждаемся, что очередь существует (или создаем)
	_, err = ch.QueueDeclare(
		queueName,
		false, // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Создаем HTTP сервер для получения данных
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var msg DataMessage
			if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
				http.Error(w, "Error decoding JSON", http.StatusBadRequest)
				return
			}
			log.Printf("Received message: %+v", msg)

			// Запись в MongoDB
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			_, err := collection.InsertOne(ctx, msg)
			cancel()
			if err != nil {
				log.Printf("Failed to insert document: %v", err)
			}

			// Отправляем сообщение в RabbitMQ
			data, err := json.Marshal(msg)
			if err != nil {
				http.Error(w, "Failed to marshal message", http.StatusInternalServerError)
				return
			}

			log.Printf("Message sent to RabbitMQ: %+v", msg)

			err = ch.Publish(
				"",        // exchange
				queueName, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        data,
				},
			)
			if err != nil {
				log.Printf("Failed to publish message to RabbitMQ: %v", err)
			}
			w.WriteHeader(http.StatusOK)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Запуск HTTP сервера
	log.Printf("Controller is listening on %s", httpServerAddr)
	log.Fatal(http.ListenAndServe(httpServerAddr, nil))
}
