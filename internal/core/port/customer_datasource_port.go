package port

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
)

type CustomerDataSource interface {
	FindByID(ctx context.Context, id uint64) (*entity.Customer, error)
	FindByCPF(ctx context.Context, cpf string) (*entity.Customer, error)
	FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]*entity.Customer, int64, error)
	Create(ctx context.Context, product *entity.Customer) error
	Update(ctx context.Context, product *entity.Customer) error
	Delete(ctx context.Context, id uint64) error
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}
