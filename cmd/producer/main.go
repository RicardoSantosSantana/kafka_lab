package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	deliveryChan := make(chan kafka.Event)
	key := []byte("minhaKey")
	producer := NewKafkaProducer()
	Publish("Mirian cortou o cabelo, e está linda!", "teste", producer, key, deliveryChan)
	go DeliveryReport(deliveryChan)
	producer.Flush(1000)

}

func NewKafkaProducer() *kafka.Producer {

	configmap := &kafka.ConfigMap{
		"bootstrap.servers":   "kafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "false",
	}

	p, err := kafka.NewProducer(configmap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}
	err := producer.Produce(message, deliveryChan)
	if err != nil {
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				/**
				* anotar no banco de dados que a mensagem foi processada.
				* ex: confirmar que uma transferência bancária ocorreu.
				 */
			}
		}
	}
}
