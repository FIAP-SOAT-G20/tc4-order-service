package entity

import (
	"time"

	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type OrderHistory struct {
	ID        uint64
	OrderID   uint64
	StaffID   *uint64
	Status    valueobject.OrderStatus
	CreatedAt time.Time
	Order     Order
	Staff     *Staff
}

func NewOrderHistory(orderID uint64, status valueobject.OrderStatus, staffID *uint64) *OrderHistory {
	return &OrderHistory{
		OrderID: orderID,
		Status:  status,
		StaffID: staffID,
	}
}
