package repository

import (
	"context"

	"library-management-system/models"
)

type BookRepository interface {
	Create(ctx context.Context, book *models.Book) error
	GetByID(ctx context.Context, id int) (*models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
	Update(ctx context.Context, book *models.Book) error
	Delete(ctx context.Context, id int) error
	Search(ctx context.Context, keyword string) ([]models.Book, error)
}
