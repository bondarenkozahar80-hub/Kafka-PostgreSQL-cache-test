package service

import "github.com/IBM/sarama"

type Service struct {
	OrderProducer *OrderProducer
}

func NewService(producer sarama.SyncProducer, kafkaTopic string) *Service {
	return &Service{
		OrderProducer: NewOrderProducer(producer, kafkaTopic),
	}
}
