package repository

import (
	"context"
	"errors"

	"library-management-system/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresMemberRepository struct {
	DB *pgxpool.Pool
}

// Convert database rows into Member models.
func scanMember(row pgx.Row) (*models.Member, error) {
	var member models.Member
	var email pgtype.Text

	err := row.Scan(
		&member.ID,
		&member.Name,
		&email,
		&member.Phone,
		&member.Address,
		&member.MembershipType,
		&member.MembershipDate,
		&member.MembershipStatus,
		&member.CreatedAt,
		&member.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if email.Valid {
		member.Email = &email.String
	}

	return &member, nil
}

// Create registers a new member.
func (r *PostgresMemberRepository) Create(ctx context.Context, member *models.Member) error {
	query := `
		INSERT INTO members
		(name, email, phone, address, membership_type,
		membership_date, membership_status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at`

	return r.DB.QueryRow(ctx, query,
		member.Name,
		member.Email,
		member.Phone,
		member.Address,
		member.MembershipType,
		member.MembershipDate,
		member.MembershipStatus,
	).Scan(&member.ID, &member.CreatedAt, &member.UpdatedAt)
}

// GetByID retrieves a member.
func (r *PostgresMemberRepository) GetByID(ctx context.Context, id int) (*models.Member, error) {
	query := `
		SELECT id, name, email, phone, address,
		membership_type, membership_date, membership_status,
		created_at, updated_at
		FROM members WHERE id = $1`

	return scanMember(r.DB.QueryRow(ctx, query, id))
}

// GetAll retrieves all members.
func (r *PostgresMemberRepository) GetAll(ctx context.Context) ([]models.Member, error) {
	query := `
		SELECT id, name, email, phone, address,
		membership_type, membership_date, membership_status,
		created_at, updated_at
		FROM members ORDER BY name ASC`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	members := make([]models.Member, 0)

	for rows.Next() {
		member, err := scanMember(rows)
		if err != nil {
			return nil, err
		}
		members = append(members, *member)
	}

	return members, rows.Err()
}

// Update member information.
func (r *PostgresMemberRepository) Update(ctx context.Context, member *models.Member) error {
	query := `
		UPDATE members SET
		name = $1, email = $2, phone = $3,
		address = $4, membership_type = $5,
		membership_status = $6,
		updated_at = CURRENT_TIMESTAMP
		WHERE id = $7`

	result, err := r.DB.Exec(ctx, query,
		member.Name,
		member.Email,
		member.Phone,
		member.Address,
		member.MembershipType,
		member.MembershipStatus,
		member.ID,
	)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

// Delete a member.
func (r *PostgresMemberRepository) Delete(ctx context.Context, id int) error {
	result, err := r.DB.Exec(ctx,
		"DELETE FROM members WHERE id = $1", id)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

// Ensure interface implementation.
var _ MemberRepository = (*PostgresMemberRepository)(nil)
