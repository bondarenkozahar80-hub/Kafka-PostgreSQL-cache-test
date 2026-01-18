package repository

import "Kafka-PostgreSQL-cache-test/internal/models"

type Orders interface {
	AddOrder(order models.Order) error
	GetOrder(OrderUID string) (*models.Order, error)
	GetOrders() ([]models.Order, error)
}
