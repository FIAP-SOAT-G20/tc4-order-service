package request

// AuthenticateBodyRequest representa o corpo da requisição de autenticação
type AuthenticateBodyRequest struct {
	CPF string `json:"cpf" binding:"required" example:"000.000.000-00"`
}
