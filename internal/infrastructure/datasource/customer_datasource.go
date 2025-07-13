package datasource

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type customerDataSource struct {
	db *gorm.DB
}

type customerKey string

func NewCustomerDataSource(db *gorm.DB) port.CustomerDataSource {
	return &customerDataSource{db}
}

func (ds *customerDataSource) FindByID(ctx context.Context, id uint64) (*entity.Customer, error) {
	var customer entity.Customer
	result := ds.db.WithContext(ctx).First(&customer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding customer: %w", result.Error)
	}
	return &customer, nil
}

func (ds *customerDataSource) FindByCPF(ctx context.Context, cpf string) (*entity.Customer, error) {
	var customer entity.Customer
	result := ds.db.WithContext(ctx).Where("cpf = ?", cpf).First(&customer)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding customer: %w", result.Error)
	}
	return &customer, nil
}

func (ds *customerDataSource) FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]*entity.Customer, int64, error) {
	var customers []*entity.Customer
	var total int64

	query := ds.db.WithContext(ctx)

	// Apply filters
	for key, value := range filters {
		switch key {
		case "name":
			if name, ok := value.(string); ok && name != "" {
				query = query.Where("name LIKE ?", "%"+name+"%")
			}
		}
	}

	// Count total before pagination
	if err := query.Model(&entity.Customer{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting customers: %w", err)
	}

	// Get paginated results
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
		return nil, 0, fmt.Errorf("error finding customers: %w", err)
	}

	return customers, total, nil
}

func (ds *customerDataSource) Create(ctx context.Context, customer *entity.Customer) error {
	if err := ds.db.WithContext(ctx).Create(customer).Error; err != nil {
		return fmt.Errorf("error creating customer: %w", err)
	}
	return nil
}

func (ds *customerDataSource) Update(ctx context.Context, customer *entity.Customer) error {
	result := ds.db.WithContext(ctx).Save(customer)
	if result.Error != nil {
		return fmt.Errorf("error updating customer: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (ds *customerDataSource) Delete(ctx context.Context, id uint64) error {
	result := ds.db.WithContext(ctx).Delete(&entity.Customer{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting customer: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (ds *customerDataSource) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return ds.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create a new context with the transaction
		keyPrincipalID := customerKey(uuid.NewString())
		txCtx := context.WithValue(ctx, keyPrincipalID, tx)
		return fn(txCtx)
	})
}
