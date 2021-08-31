package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func InitRabbitMq(config RabbitMQConfig) (*amqp.Connection, *amqp.Channel) {
	host := config.Host
	port := config.Port
	username := config.UserName
	password := config.Password
	vHost := config.Vhost
	rabbitmqConnString := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", username, password, host, port, vHost)
	conn, ch := connect(rabbitmqConnString, config.Exchange, config.Durable)
	return conn, ch
}

func connect(url string, exchange string, durable bool) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "[ERROR] FAILED TO CONNECT TO RABBITMQ")

	ch, err := conn.Channel()
	failOnError(err, "[ERROR] FAILED TO OPEN A CHANNEL RABBITMQ")

	_ = ch.ExchangeDeclare(
		exchange,
		"topic",
		durable,  // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	return conn, ch
}

func DeclareQueue(ch *amqp.Channel, name string, exchange string, durable bool) amqp.Queue {
	q, err := ch.QueueDeclare(
		name, // name
		durable,                         // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
	failOnError(err, "[ERROR] FAILED TO DECLARE A QUEUE")
	err = ch.QueueBind(
		q.Name,                                 // queue name
		fmt.Sprintf("%s.%s", exchange, q.Name), // routing key
		exchange,                               // exchange
		false,
		nil,
	)
	failOnError(err, "[ERROR] FAILED TO BINDING A QUEUE")
	log.Printf("[INFO] DECLARED QUEUE " + name)
	return q
}

func Publish(ch *amqp.Channel, q amqp.Queue, data []byte)  {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", q.Name)
}

func Consume(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {
	log.Printf("[INFO] CONSUMING " + q.Name)
	messages, err := ch.Consume(
		q.Name, // queue
		q.Name, // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "[ERROR] FAILED TO REGISTER A CONSUMER")
	return messages
}

func HandleMessages(qName string, messages <-chan amqp.Delivery, f func(d []byte)) {
	forever := make(chan bool)

	go func() {
		for d := range messages {
			fmt.Printf("[INFO] RECEIVED A MESSAGE OF %s: %s", qName, d.Body)
			f(d.Body)
		}
	}()

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}