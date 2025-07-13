package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
)

type StaffGateway interface {
	FindByID(ctx context.Context, id uint64) (*entity.Staff, error)
	FindAll(ctx context.Context, name string, role valueobject.StaffRole, page, limit int) ([]*entity.Staff, int64, error)
	Create(ctx context.Context, staff *entity.Staff) error
	Update(ctx context.Context, staff *entity.Staff) error
	Delete(ctx context.Context, id uint64) error
}
