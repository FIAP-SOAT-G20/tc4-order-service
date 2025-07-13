package entity

import (
	"time"

	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type Order struct {
	ID            uint64
	CustomerID    uint64
	Status        valueobject.OrderStatus
	Payment       Payment
	Customer      Customer
	OrderProducts []OrderProduct
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (p *Order) Update(customerID uint64, status valueobject.OrderStatus) {
	if customerID != 0 {
		p.CustomerID = customerID
	}
	if status != valueobject.UNDEFINDED {
		p.Status = status
	}
	p.OrderProducts = nil
	p.UpdatedAt = time.Now()
}
