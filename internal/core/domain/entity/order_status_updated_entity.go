package entity

type OrderStatusUpdated struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
	StaffID *string `json:"staff_id,omitempty"` // Optional field for staff ID
}


