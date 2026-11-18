package repository

import (
	"context"
	"errors"
	"fmt"

	"library-management-system/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresLoanRepository struct {
	DB *pgxpool.Pool
}

// IssueBook creates a loan and decreases book availability.
func (r *PostgresLoanRepository) IssueBook(
	ctx context.Context,
	loan *models.Loan,
) error {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Lock the book row to prevent simultaneous over-issuing.
	var available int
	err = tx.QueryRow(ctx,
		`SELECT available_copies FROM books WHERE id = $1 FOR UPDATE`,
		loan.BookID,
	).Scan(&available)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	if available <= 0 {
		return errors.New("book is not available")
	}

	// Check that the member exists and is active.
	var memberStatus string
	err = tx.QueryRow(ctx,
		`SELECT membership_status FROM members WHERE id = $1 FOR UPDATE`,
		loan.MemberID,
	).Scan(&memberStatus)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	if memberStatus != "Active" {
		return errors.New("member is not active")
	}

	// Insert the loan record.
	err = tx.QueryRow(ctx, `
		INSERT INTO loans (book_id, member_id, issue_date, due_date, status)
		VALUES ($1, $2, $3, $4, 'Issued')
		RETURNING id, created_at
	`,
		loan.BookID,
		loan.MemberID,
		loan.IssueDate,
		loan.DueDate,
	).Scan(&loan.ID, &loan.CreatedAt)

	if err != nil {
		return err
	}

	// Decrease available copies.
	_, err = tx.Exec(ctx, `
		UPDATE books
		SET available_copies = available_copies - 1,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`, loan.BookID)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// ReturnBook processes a book return.
func (r *PostgresLoanRepository) ReturnBook(
	ctx context.Context,
	loanID int,
) error {

	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	var bookID int
	var status string

	// Lock the loan to prevent duplicate returns.
	err = tx.QueryRow(ctx, `
		SELECT book_id, status
		FROM loans
		WHERE id = $1
		FOR UPDATE
	`, loanID).Scan(&bookID, &status)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}

	if status == "Returned" {
		return errors.New("book has already been returned")
	}

	// Update the loan record.
	_, err = tx.Exec(ctx, `
		UPDATE loans
		SET status = 'Returned',
		    returned_at = CURRENT_TIMESTAMP
		WHERE id = $1
	`, loanID)

	if err != nil {
		return err
	}

	// Restore book availability.
	result, err := tx.Exec(ctx, `
		UPDATE books
		SET available_copies = available_copies + 1,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $1
		  AND available_copies < total_copies
	`, bookID)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("book availability could not be restored")
	}

	return tx.Commit(ctx)
}

// GetAll retrieves all loans.
func (r *PostgresLoanRepository) GetAll(
	ctx context.Context,
) ([]models.Loan, error) {

	return r.getLoans(ctx, `
		SELECT id, book_id, member_id, issue_date,
		       due_date, returned_at, status, created_at
		FROM loans
		ORDER BY issue_date DESC
	`)
}

// GetOverdue retrieves currently overdue loans.
func (r *PostgresLoanRepository) GetOverdue(
	ctx context.Context,
) ([]models.Loan, error) {

	return r.getLoans(ctx, `
		SELECT id, book_id, member_id, issue_date,
		       due_date, returned_at, status, created_at
		FROM loans
		WHERE status = 'Issued'
		  AND due_date < CURRENT_DATE
		ORDER BY due_date ASC
	`)
}

// GetHistoryByMember retrieves a member's borrowing history.
func (r *PostgresLoanRepository) GetHistoryByMember(
	ctx context.Context,
	memberID int,
) ([]models.Loan, error) {

	return r.getLoans(ctx, `
		SELECT id, book_id, member_id, issue_date,
		       due_date, returned_at, status, created_at
		FROM loans
		WHERE member_id = $1
		ORDER BY issue_date DESC
	`, memberID)
}

// GetHistoryByBook retrieves a book's borrowing history.
func (r *PostgresLoanRepository) GetHistoryByBook(
	ctx context.Context,
	bookID int,
) ([]models.Loan, error) {

	return r.getLoans(ctx, `
		SELECT id, book_id, member_id, issue_date,
		       due_date, returned_at, status, created_at
		FROM loans
		WHERE book_id = $1
		ORDER BY issue_date DESC
	`, bookID)
}

// Shared helper for retrieving loan records.
func (r *PostgresLoanRepository) getLoans(
	ctx context.Context,
	query string,
	args ...interface{},
) ([]models.Loan, error) {

	rows, err := r.DB.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	loans := make([]models.Loan, 0)

	for rows.Next() {
		var loan models.Loan

		err := rows.Scan(
			&loan.ID,
			&loan.BookID,
			&loan.MemberID,
			&loan.IssueDate,
			&loan.DueDate,
			&loan.ReturnedAt,
			&loan.Status,
			&loan.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		loans = append(loans, loan)
	}

	return loans, rows.Err()
}

// Verify interface implementation.
var _ LoanRepository = (*PostgresLoanRepository)(nil)
