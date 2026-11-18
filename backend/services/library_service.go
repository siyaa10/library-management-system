package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"library-management-system/models"
	"library-management-system/repository"

	"github.com/jackc/pgx/v5/pgconn"
)

type LibraryService struct {
	Books   repository.BookRepository
	Members repository.MemberRepository
	Loans   repository.LoanRepository
}

// Constructor for the service.
func NewLibraryService(
	books repository.BookRepository,
	members repository.MemberRepository,
	loans repository.LoanRepository,
) *LibraryService {
	return &LibraryService{
		Books:   books,
		Members: members,
		Loans:   loans,
	}
}

// Add a new book after validation.
func (s *LibraryService) AddBook(ctx context.Context, book *models.Book) error {
	book.Title = strings.TrimSpace(book.Title)
	book.Author = strings.TrimSpace(book.Author)
	book.Category = strings.TrimSpace(book.Category)
	book.ISBN = strings.TrimSpace(book.ISBN)

	if book.Title == "" || book.Author == "" ||
		book.Category == "" || book.ISBN == "" {
		return errors.New("title, author, category and ISBN are required")
	}

	if book.TotalCopies < 1 {
		return errors.New("total copies must be at least one")
	}

	book.AvailableCopies = book.TotalCopies

	return s.Books.Create(ctx, book)
}

// Retrieve all books.
func (s *LibraryService) GetBooks(ctx context.Context) ([]models.Book, error) {
	return s.Books.GetAll(ctx)
}

// Retrieve one book.
func (s *LibraryService) GetBook(ctx context.Context, id int) (*models.Book, error) {
	if id <= 0 {
		return nil, errors.New("invalid book ID")
	}
	return s.Books.GetByID(ctx, id)
}

// Update a book.
func (s *LibraryService) UpdateBook(ctx context.Context, book *models.Book) error {
	if book.ID <= 0 || strings.TrimSpace(book.Title) == "" ||
		strings.TrimSpace(book.Author) == "" {
		return errors.New("valid book ID, title and author are required")
	}

	existing, err := s.Books.GetByID(ctx, book.ID)
	if err != nil {
		return err
	}

	issuedCopies := existing.TotalCopies - existing.AvailableCopies

	if book.TotalCopies < issuedCopies || book.TotalCopies < 1 {
		return errors.New("total copies cannot be less than currently issued copies")
	}

	book.AvailableCopies = book.TotalCopies - issuedCopies

	return s.Books.Update(ctx, book)
}

// Delete a book.
func (s *LibraryService) DeleteBook(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid book ID")
	}

	book, err := s.Books.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if book.AvailableCopies != book.TotalCopies {
		return errors.New("cannot delete a book with active loans")
	}

	err = s.Books.Delete(ctx, id)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return errors.New("This book cannot be deleted because borrowing history exists.")
		}

		return err
	}

	return nil
}

// Search books.
func (s *LibraryService) SearchBooks(ctx context.Context, keyword string) ([]models.Book, error) {
	return s.Books.Search(ctx, strings.TrimSpace(keyword))
}

// Register a member.
func (s *LibraryService) AddMember(ctx context.Context, member *models.Member) error {
	member.Name = strings.TrimSpace(member.Name)
	member.Phone = strings.TrimSpace(member.Phone)

	if member.Name == "" || member.Phone == "" {
		return errors.New("member name and phone are required")
	}

	if member.MembershipType == "" {
		member.MembershipType = "Regular"
	}

	if member.MembershipDate.IsZero() {
		member.MembershipDate = time.Now()
	}

	member.MembershipStatus = "Active"

	return s.Members.Create(ctx, member)
}

// Retrieve all members.
func (s *LibraryService) GetMembers(ctx context.Context) ([]models.Member, error) {
	return s.Members.GetAll(ctx)
}

// Update a member.
func (s *LibraryService) UpdateMember(ctx context.Context, member *models.Member) error {
	if member.ID <= 0 || strings.TrimSpace(member.Name) == "" ||
		strings.TrimSpace(member.Phone) == "" {
		return errors.New("valid member ID, name and phone are required")
	}

	return s.Members.Update(ctx, member)
}

// Delete a member.

// Delete a member.
func (s *LibraryService) DeleteMember(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("invalid member ID")
	}

	err := s.Members.Delete(ctx, id)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return errors.New("This member cannot be deleted because borrowing history exists.")
		}

		return err
	}

	return nil
}

// Issue a book to a member.
func (s *LibraryService) IssueBook(
	ctx context.Context,
	bookID int,
	memberID int,
	dueDate time.Time,
) error {

	if bookID <= 0 || memberID <= 0 {
		return errors.New("valid book and member IDs are required")
	}

	today := time.Now()
	due := dueDate

	if due.Before(time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())) {
		return errors.New("due date cannot be in the past")
	}

	loan := &models.Loan{
		BookID:    bookID,
		MemberID:  memberID,
		IssueDate: today,
		DueDate:   due,
		Status:    "Issued",
	}

	return s.Loans.IssueBook(ctx, loan)
}

// Return an issued book.
func (s *LibraryService) ReturnBook(ctx context.Context, loanID int) error {
	if loanID <= 0 {
		return errors.New("invalid loan ID")
	}

	return s.Loans.ReturnBook(ctx, loanID)
}

// Retrieve overdue loans.
func (s *LibraryService) GetOverdueLoans(ctx context.Context) ([]models.Loan, error) {
	return s.Loans.GetOverdue(ctx)
}

// Retrieve member borrowing history.
func (s *LibraryService) GetMemberHistory(ctx context.Context, memberID int) ([]models.Loan, error) {
	if memberID <= 0 {
		return nil, errors.New("invalid member ID")
	}

	return s.Loans.GetHistoryByMember(ctx, memberID)
}

// Retrieve book borrowing history.
func (s *LibraryService) GetBookHistory(ctx context.Context, bookID int) ([]models.Loan, error) {
	if bookID <= 0 {
		return nil, errors.New("invalid book ID")
	}

	return s.Loans.GetHistoryByBook(ctx, bookID)
}

// Retrieve one member by ID.
func (s *LibraryService) GetMember(ctx context.Context, id int) (*models.Member, error) {
	if id <= 0 {
		return nil, errors.New("invalid member ID")
	}
	return s.Members.GetByID(ctx, id)
}

// Retrieve all loan records.
func (s *LibraryService) GetLoans(ctx context.Context) ([]models.Loan, error) {
	return s.Loans.GetAll(ctx)
}
