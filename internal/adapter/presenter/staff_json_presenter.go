package presenter

import (
	"encoding/json"
	"errors"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/dto"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type staffJsonPresenter struct{}

// StaffJsonResponse represents the response of a staff
func NewStaffJsonPresenter() port.Presenter {
	return &staffJsonPresenter{}
}

// toStaffJsonResponse convert entity.Staff to StaffJsonResponse
func toStaffJsonResponse(staff *entity.Staff) StaffJsonResponse {
	return StaffJsonResponse{
		ID:        staff.ID,
		Name:      staff.Name,
		Role:      staff.Role.String(),
		CreatedAt: staff.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: staff.UpdatedAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
	}
}

// Present write the response to the client
func (p *staffJsonPresenter) Present(pp dto.PresenterInput) ([]byte, error) {
	switch v := pp.Result.(type) {
	case *entity.Staff:
		output := toStaffJsonResponse(v)
		return json.Marshal(output)
	case []*entity.Staff:
		staffOutputs := make([]StaffJsonResponse, len(v))
		for i, staff := range v {
			staffOutputs[i] = toStaffJsonResponse(staff)
		}

		output := &StaffJsonPaginatedResponse{
			JsonPagination: JsonPagination{
				Total: pp.Total,
				Page:  pp.Page,
				Limit: pp.Limit,
			},
			Staffs: staffOutputs,
		}
		return json.Marshal(output)
	default:
		return nil, domain.NewInternalError(errors.New(domain.ErrInternalError))
	}
}
