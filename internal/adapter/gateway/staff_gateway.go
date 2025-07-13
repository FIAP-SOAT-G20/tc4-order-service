package gateway

import (
	"context"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	valueobject "github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/value_object"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type staffGateway struct {
	dataSource port.StaffDataSource
}

func NewStaffGateway(dataSource port.StaffDataSource) port.StaffGateway {
	return &staffGateway{dataSource}
}

func (g *staffGateway) FindByID(ctx context.Context, id uint64) (*entity.Staff, error) {
	return g.dataSource.FindByID(ctx, id)
}

func (g *staffGateway) FindAll(ctx context.Context, name string, role valueobject.StaffRole, page, limit int) ([]*entity.Staff, int64, error) {
	filters := make(map[string]interface{})

	if name != "" {
		filters["name"] = name
	}
	if role != "" {
		filters["role"] = role.String()
	}

	return g.dataSource.FindAll(ctx, filters, page, limit)
}

func (g *staffGateway) Create(ctx context.Context, product *entity.Staff) error {
	return g.dataSource.Create(ctx, product)
}

func (g *staffGateway) Update(ctx context.Context, product *entity.Staff) error {
	return g.dataSource.Update(ctx, product)
}

func (g *staffGateway) Delete(ctx context.Context, id uint64) error {
	return g.dataSource.Delete(ctx, id)
}
