package datasource

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/domain/entity"
	"github.com/FIAP-SOAT-G20/fiap-tech-challenge-3-api/internal/core/port"
)

type staffDataSource struct {
	db *gorm.DB
}

type staffKey string

func NewStaffDataSource(db *gorm.DB) port.StaffDataSource {
	return &staffDataSource{db}
}

func (ds *staffDataSource) FindByID(ctx context.Context, id uint64) (*entity.Staff, error) {
	var staff entity.Staff
	result := ds.db.WithContext(ctx).First(&staff, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error finding staff: %w", result.Error)
	}
	return &staff, nil
}

func (ds *staffDataSource) FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]*entity.Staff, int64, error) {
	var staffs []*entity.Staff
	var total int64

	query := ds.db.WithContext(ctx)

	// Apply filters
	for key, value := range filters {
		switch key {
		case "name":
			if name, ok := value.(string); ok && name != "" {
				query = query.Where("name LIKE ?", "%"+name+"%")
			}
		case "role":
			if role, ok := value.(string); ok && role != "" {
				query = query.Where("role LIKE ?", "%"+role+"%")
			}
		}

	}

	// Count total before pagination
	if err := query.Model(&entity.Staff{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting staffs: %w", err)
	}

	// Get paginated results
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&staffs).Error; err != nil {
		return nil, 0, fmt.Errorf("error finding staffs: %w", err)
	}

	return staffs, total, nil
}

func (ds *staffDataSource) Create(ctx context.Context, staff *entity.Staff) error {
	if err := ds.db.WithContext(ctx).Create(staff).Error; err != nil {
		return fmt.Errorf("error creating staff: %w", err)
	}
	return nil
}

func (ds *staffDataSource) Update(ctx context.Context, staff *entity.Staff) error {
	result := ds.db.WithContext(ctx).Save(staff)
	if result.Error != nil {
		return fmt.Errorf("error updating staff: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (ds *staffDataSource) Delete(ctx context.Context, id uint64) error {
	result := ds.db.WithContext(ctx).Delete(&entity.Staff{}, id)
	if result.Error != nil {
		return fmt.Errorf("error deleting staff: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return nil
	}
	return nil
}

func (ds *staffDataSource) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return ds.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Create a new context with the transaction
		keyPrincipalID := staffKey(uuid.NewString())
		txCtx := context.WithValue(ctx, keyPrincipalID, tx)
		return fn(txCtx)
	})
}
