package initializers

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var ChannelAmqp *amqp.Channel

func ConnectToRabbitMQ() {
   amqpConnection, err := amqp.Dial(os.Getenv(
      "RABBITMQ_URI"))
   if err != nil {
       log.Fatal(err, "Failed to connect to MQ")
   }
   ChannelAmqp, _ = amqpConnection.Channel()
}