package repository

import (
	"context"
	"errors"

	"library-management-system/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("record not found")

type PostgresBookRepository struct {
	DB *pgxpool.Pool
}

// Helper function to scan database rows into a Book model.
func scanBook(row pgx.Row) (*models.Book, error) {
	var book models.Book
	var description pgtype.Text

	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Category,
		&book.ISBN,
		&book.TotalCopies,
		&book.AvailableCopies,
		&description,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if description.Valid {
		book.Description = &description.String
	}

	return &book, nil
}

// Create adds a new book.
func (r *PostgresBookRepository) Create(ctx context.Context, book *models.Book) error {
	query := `
		INSERT INTO books
		(title, author, category, isbn, total_copies, available_copies, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`

	return r.DB.QueryRow(ctx, query,
		book.Title,
		book.Author,
		book.Category,
		book.ISBN,
		book.TotalCopies,
		book.AvailableCopies,
		book.Description,
	).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)
}

// GetByID retrieves one book.
func (r *PostgresBookRepository) GetByID(ctx context.Context, id int) (*models.Book, error) {
	query := `
		SELECT id, title, author, category, isbn,
		total_copies, available_copies, description,
		created_at, updated_at
		FROM books WHERE id = $1`

	return scanBook(r.DB.QueryRow(ctx, query, id))
}

// GetAll retrieves all books.
func (r *PostgresBookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	query := `
		SELECT id, title, author, category, isbn,
		total_copies, available_copies, description,
		created_at, updated_at
		FROM books ORDER BY title ASC`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]models.Book, 0)

	for rows.Next() {
		book, err := scanBook(rows)
		if err != nil {
			return nil, err
		}
		books = append(books, *book)
	}

	return books, rows.Err()
}

// Update modifies an existing book.
func (r *PostgresBookRepository) Update(ctx context.Context, book *models.Book) error {
	query := `
		UPDATE books SET
		title = $1, author = $2, category = $3,
		isbn = $4, total_copies = $5,
		available_copies = $6, description = $7,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = $8`

	result, err := r.DB.Exec(ctx, query,
		book.Title,
		book.Author,
		book.Category,
		book.ISBN,
		book.TotalCopies,
		book.AvailableCopies,
		book.Description,
		book.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

// Delete removes a book.
func (r *PostgresBookRepository) Delete(ctx context.Context, id int) error {
	result, err := r.DB.Exec(ctx, "DELETE FROM books WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

// Search finds books by title, author, ISBN or category.
func (r *PostgresBookRepository) Search(ctx context.Context, keyword string) ([]models.Book, error) {
	query := `
		SELECT id, title, author, category, isbn,
		total_copies, available_copies, description,
		created_at, updated_at
		FROM books
		WHERE title ILIKE $1
		OR author ILIKE $1
		OR isbn ILIKE $1
		OR category ILIKE $1
		ORDER BY title ASC`

	rows, err := r.DB.Query(ctx, query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]models.Book, 0)

	for rows.Next() {
		book, err := scanBook(rows)
		if err != nil {
			return nil, err
		}
		books = append(books, *book)
	}

	return books, rows.Err()
}

// Ensure the implementation satisfies the interface.
var _ BookRepository = (*PostgresBookRepository)(nil)
