package dto

type CreatePaymentInput struct {
	OrderID uint64
}

type UpdatePaymentInput struct {
	Resource string
	Topic    string
}

type GetPaymentInput struct {
	OrderID uint64
}
