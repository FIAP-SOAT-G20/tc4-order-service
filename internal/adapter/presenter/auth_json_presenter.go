package presenter

import (
	"errors"

	"encoding/json"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type authPresenter struct{}

func NewAuthPresenter() port.Presenter {
	return &authPresenter{}
}

func ToTokenResponse(token string) AuthenticationResponse {
	return AuthenticationResponse{
		AccessToken: token,
	}
}

func (p *authPresenter) Present(input dto.PresenterInput) ([]byte, error) {
	switch v := input.Result.(type) {
	case string:
		output := ToTokenResponse(v)
		return json.Marshal(output)
	default:
		return nil, domain.NewInternalError(errors.New(domain.ErrInternalError))
	}
}
