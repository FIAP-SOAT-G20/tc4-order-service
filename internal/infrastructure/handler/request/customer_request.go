package request

type ListCustomersQueryRequest struct {
	Name  string `form:"name" example:"John Doe"`
	Page  int    `form:"page,default=1" example:"1"`
	Limit int    `form:"limit,default=10" example:"10"`
}

type CreateCustomerBodyRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100" example:"John Doe"`
	Email string `json:"email" binding:"required,email" example:"john.doe@email.com"`
	CPF   string `json:"cpf" binding:"required" example:"123.456.789-00"`
}

type UpdateCustomerUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type UpdateCustomerBodyRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100" example:"Produto A"`
	Email string `json:"email" binding:"required,email" example:"test.customer.1@email.com"`
}

type GetCustomerUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

type DeleteCustomerUriRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}
