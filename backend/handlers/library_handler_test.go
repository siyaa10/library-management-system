package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"library-management-system/models"
	"library-management-system/repository"
	"library-management-system/services"
)

type testBookRepository struct{}

func (r *testBookRepository) Create(ctx context.Context, book *models.Book) error {
	return nil
}

func (r *testBookRepository) GetByID(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}

func (r *testBookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	return []models.Book{}, nil
}

func (r *testBookRepository) Update(ctx context.Context, book *models.Book) error {
	return nil
}

func (r *testBookRepository) Delete(ctx context.Context, id int) error {
	return nil
}

func (r *testBookRepository) Search(ctx context.Context, keyword string) ([]models.Book, error) {
	return []models.Book{}, nil
}

func createTestHandler() *LibraryHandler {
	var books repository.BookRepository = &testBookRepository{}
	service := services.NewLibraryService(books, nil, nil)
	return NewLibraryHandler(service, nil)
}

func TestGetBooksAPI(t *testing.T) {
	handler := createTestHandler()

	request := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	response := httptest.NewRecorder()

	handler.GetBooks(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", response.Code)
	}
}

func TestCreateBookInvalidData(t *testing.T) {
	handler := createTestHandler()

	request := httptest.NewRequest(
		http.MethodPost,
		"/api/books",
		strings.NewReader("{invalid json"),
	)
	response := httptest.NewRecorder()

	handler.CreateBook(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("Expected status 400, got %d", response.Code)
	}
}
