package services

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	"library-management-system/models"
)

func TestAddBookValidation(t *testing.T) {
	service := NewLibraryService(nil, nil, nil)

	book := &models.Book{
		Title:       "",
		Author:      "James Clear",
		Category:    "Self Help",
		ISBN:        "9780735211292",
		TotalCopies: 5,
	}

	err := service.AddBook(context.Background(), book)

	if err == nil {
		t.Fatal("Expected an error when book title is empty")
	}

	if !strings.Contains(err.Error(), "required") {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestAddBookInvalidCopies(t *testing.T) {
	service := NewLibraryService(nil, nil, nil)

	book := &models.Book{
		Title:       "Atomic Habits",
		Author:      "James Clear",
		Category:    "Self Help",
		ISBN:        "9780735211292",
		TotalCopies: 0,
	}

	err := service.AddBook(context.Background(), book)

	if err == nil {
		t.Fatal("Expected an error when total copies are zero")
	}
}

func TestAddMemberValidation(t *testing.T) {
	service := NewLibraryService(nil, nil, nil)

	member := &models.Member{
		Name:  "",
		Phone: "9876543210",
	}

	err := service.AddMember(context.Background(), member)

	if err == nil {
		t.Fatal("Expected an error when member name is empty")
	}
}

type mockBookRepository struct {
	createdBook *models.Book
	book        *models.Book
	deletedID   int
	updatedBook *models.Book
}

func (m *mockBookRepository) Create(ctx context.Context, book *models.Book) error {
	m.createdBook = book
	return nil
}

func (m *mockBookRepository) GetByID(ctx context.Context, id int) (*models.Book, error) {
	return m.book, nil
}

func (m *mockBookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	return nil, nil
}

func (m *mockBookRepository) Update(ctx context.Context, book *models.Book) error {
	m.updatedBook = book
	return nil
}

func (m *mockBookRepository) Delete(ctx context.Context, id int) error {
	m.deletedID = id
	return nil
}

func (m *mockBookRepository) Search(ctx context.Context, keyword string) ([]models.Book, error) {
	return nil, nil
}

func TestAddBookSuccess(t *testing.T) {
	mockRepo := &mockBookRepository{}
	service := NewLibraryService(mockRepo, nil, nil)

	book := &models.Book{
		Title:       "Atomic Habits",
		Author:      "James Clear",
		Category:    "Self Help",
		ISBN:        "9780735211292",
		TotalCopies: 5,
	}

	err := service.AddBook(context.Background(), book)

	if err != nil {
		t.Fatalf("Expected book to be added successfully, got: %v", err)
	}

	if mockRepo.createdBook.AvailableCopies != 5 {
		t.Errorf("Expected 5 available copies, got %d", mockRepo.createdBook.AvailableCopies)
	}
}
func TestDeleteBookWithActiveLoans(t *testing.T) {
	mockRepo := &mockBookRepository{
		book: &models.Book{
			ID:              1,
			TotalCopies:     5,
			AvailableCopies: 3,
		},
	}

	service := NewLibraryService(mockRepo, nil, nil)

	err := service.DeleteBook(context.Background(), 1)

	if err == nil {
		t.Fatal("Expected error when deleting a book with active loans")
	}

	if mockRepo.deletedID != 0 {
		t.Fatal("Book should not have been deleted")
	}
}

func TestDeleteBookSuccess(t *testing.T) {
	mockRepo := &mockBookRepository{
		book: &models.Book{
			ID:              2,
			TotalCopies:     5,
			AvailableCopies: 5,
		},
	}

	service := NewLibraryService(mockRepo, nil, nil)

	err := service.DeleteBook(context.Background(), 2)

	if err != nil {
		t.Fatalf("Expected successful deletion, got: %v", err)
	}

	if mockRepo.deletedID != 2 {
		t.Errorf("Expected deleted ID 2, got %d", mockRepo.deletedID)
	}
}

func TestUpdateBookWithTooFewCopies(t *testing.T) {
	mockRepo := &mockBookRepository{
		book: &models.Book{
			ID:              1,
			TotalCopies:     5,
			AvailableCopies: 3,
		},
	}

	service := NewLibraryService(mockRepo, nil, nil)

	book := &models.Book{
		ID:          1,
		Title:       "Atomic Habits",
		Author:      "James Clear",
		TotalCopies: 1,
	}

	err := service.UpdateBook(context.Background(), book)

	if err == nil {
		t.Fatal("Expected error when reducing copies below issued copies")
	}
}

func TestUpdateBookSuccess(t *testing.T) {
	mockRepo := &mockBookRepository{
		book: &models.Book{
			ID:              1,
			TotalCopies:     5,
			AvailableCopies: 3,
		},
	}

	service := NewLibraryService(mockRepo, nil, nil)

	book := &models.Book{
		ID:          1,
		Title:       "Atomic Habits",
		Author:      "James Clear",
		TotalCopies: 6,
	}

	err := service.UpdateBook(context.Background(), book)

	if err != nil {
		t.Fatalf("Expected successful update, got: %v", err)
	}

	if book.AvailableCopies != 4 {
		t.Errorf("Expected 4 available copies, got %d", book.AvailableCopies)
	}
}

// Thread-safe mock repository for concurrency testing.
type concurrentBookRepository struct {
	mu    sync.Mutex
	books []models.Book
}

func (m *concurrentBookRepository) Create(ctx context.Context, book *models.Book) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.books = append(m.books, *book)
	return nil
}

func (m *concurrentBookRepository) GetByID(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}

func (m *concurrentBookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	return nil, nil
}

func (m *concurrentBookRepository) Update(ctx context.Context, book *models.Book) error {
	return nil
}

func (m *concurrentBookRepository) Delete(ctx context.Context, id int) error {
	return nil
}

func (m *concurrentBookRepository) Search(ctx context.Context, keyword string) ([]models.Book, error) {
	return nil, nil
}

func TestConcurrentBookCreation(t *testing.T) {
	mockRepo := &concurrentBookRepository{
		books: make([]models.Book, 0),
	}

	service := NewLibraryService(mockRepo, nil, nil)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			book := &models.Book{
				Title:       "Test Book",
				Author:      "Test Author",
				Category:    "Testing",
				ISBN:        fmt.Sprintf("TEST-%d", i),
				TotalCopies: 1,
			}

			if err := service.AddBook(context.Background(), book); err != nil {
				t.Errorf("Failed to add book: %v", err)
			}
		}(i)
	}

	wg.Wait()

	mockRepo.mu.Lock()
	defer mockRepo.mu.Unlock()

	if len(mockRepo.books) != 10 {
		t.Errorf("Expected 10 books, got %d", len(mockRepo.books))
	}
}
