package repository

import (
	"context"

	"library-management-system/models"
)

type LoanRepository interface {
	IssueBook(ctx context.Context, loan *models.Loan) error
	ReturnBook(ctx context.Context, loanID int) error
	GetAll(ctx context.Context) ([]models.Loan, error)
	GetOverdue(ctx context.Context) ([]models.Loan, error)
	GetHistoryByMember(ctx context.Context, memberID int) ([]models.Loan, error)
	GetHistoryByBook(ctx context.Context, bookID int) ([]models.Loan, error)
}
