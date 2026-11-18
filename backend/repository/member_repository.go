package repository

import (
	"context"

	"library-management-system/models"
)

type MemberRepository interface {
	Create(ctx context.Context, member *models.Member) error
	GetByID(ctx context.Context, id int) (*models.Member, error)
	GetAll(ctx context.Context) ([]models.Member, error)
	Update(ctx context.Context, member *models.Member) error
	Delete(ctx context.Context, id int) error
}
