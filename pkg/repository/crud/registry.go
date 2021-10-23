package crud

import (
	"context"

	"github.com/bm-krishna/tenant-service/pkg/models"
)

// TenantRepo Interface provide api for repository
type TenantRepo interface {
	Create(ctx context.Context, tenant *models.Tenant) (int64, error)
	Fetch(ctx context.Context, count int64) ([]*models.Tenant, error)
	GetByID(ctx context.Context, tenantID int64) (*models.Tenant, error)
	Update(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error)
	Delete(ctx context.Context, tenantID int64) (bool, error)
}
