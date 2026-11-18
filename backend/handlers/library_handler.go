package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"library-management-system/models"
	"library-management-system/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LibraryHandler struct {
	Service *services.LibraryService
	DB      *pgxpool.Pool
}

func NewLibraryHandler(service *services.LibraryService, db *pgxpool.Pool) *LibraryHandler {
	return &LibraryHandler{
		Service: service,
		DB:      db,
	}
}

// Send JSON responses.
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Send error responses.
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// Decode incoming JSON.
func decodeJSON(r *http.Request, destination interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(destination)
}

// Parse an ID from the URL.
func parseID(r *http.Request) (int, error) {
	return strconv.Atoi(r.PathValue("id"))
}

// Register all API routes.
func (h *LibraryHandler) RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/dashboard", h.GetDashboard)

	mux.HandleFunc("GET /api/books", h.GetBooks)
	mux.HandleFunc("POST /api/books", h.CreateBook)
	mux.HandleFunc("GET /api/books/{id}", h.GetBook)
	mux.HandleFunc("PUT /api/books/{id}", h.UpdateBook)
	mux.HandleFunc("DELETE /api/books/{id}", h.DeleteBook)

	mux.HandleFunc("GET /api/members", h.GetMembers)
	mux.HandleFunc("POST /api/members", h.CreateMember)
	mux.HandleFunc("GET /api/members/{id}", h.GetMember)
	mux.HandleFunc("PUT /api/members/{id}", h.UpdateMember)
	mux.HandleFunc("DELETE /api/members/{id}", h.DeleteMember)

	mux.HandleFunc("GET /api/loans", h.GetLoans)
	mux.HandleFunc("POST /api/loans", h.IssueBook)
	mux.HandleFunc("PUT /api/loans/{id}/return", h.ReturnBook)
	mux.HandleFunc("GET /api/loans/overdue", h.GetOverdue)
	mux.HandleFunc("GET /api/members/{id}/history", h.GetMemberHistory)
	mux.HandleFunc("GET /api/books/{id}/history", h.GetBookHistory)
}

// BOOK HANDLERS

func (h *LibraryHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	keyword := strings.TrimSpace(r.URL.Query().Get("search"))

	var books []models.Book
	var err error

	if keyword != "" {
		books, err = h.Service.SearchBooks(ctx, keyword)
	} else {
		books, err = h.Service.GetBooks(ctx)
	}

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Filter books by category.
	category := strings.TrimSpace(r.URL.Query().Get("category"))

	if category != "" {
		filteredBooks := make([]models.Book, 0)

		for _, book := range books {
			if strings.EqualFold(book.Category, category) {
				filteredBooks = append(filteredBooks, book)
			}
		}

		books = filteredBooks
	}

	// Sort books by title, author or ISBN.
	sortBy := strings.ToLower(r.URL.Query().Get("sortBy"))
	sortOrder := strings.ToLower(r.URL.Query().Get("order"))

	if sortBy != "" {
		if sortBy != "title" && sortBy != "author" && sortBy != "isbn" {
			respondError(w, http.StatusBadRequest, "Invalid sort field")
			return
		}

		sort.Slice(books, func(i, j int) bool {
			var a, b string

			switch sortBy {
			case "title":
				a, b = books[i].Title, books[j].Title
			case "author":
				a, b = books[i].Author, books[j].Author
			case "isbn":
				a, b = books[i].ISBN, books[j].ISBN
			}

			if sortOrder == "desc" {
				return strings.ToLower(a) > strings.ToLower(b)
			}

			return strings.ToLower(a) < strings.ToLower(b)
		})
	}

	respondJSON(w, http.StatusOK, books)
}

func (h *LibraryHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := decodeJSON(r, &book); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book data")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.Service.AddBook(ctx, &book); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, book)
}

func (h *LibraryHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	book, err := h.Service.GetBook(ctx, id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, book)
}

func (h *LibraryHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book models.Book
	if err := decodeJSON(r, &book); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book data")
		return
	}

	book.ID = id

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.Service.UpdateBook(ctx, &book); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, book)
}

func (h *LibraryHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.Service.DeleteBook(ctx, id); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}

// MEMBER HANDLERS

func (h *LibraryHandler) GetMembers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	members, err := h.Service.GetMembers(ctx)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, members)
}

func (h *LibraryHandler) CreateMember(w http.ResponseWriter, r *http.Request) {
	var member models.Member

	if err := decodeJSON(r, &member); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid member data")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.Service.AddMember(ctx, &member); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, member)
}

func (h *LibraryHandler) GetMember(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid member ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	member, err := h.Service.GetMember(ctx, id)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, member)
}

func (h *LibraryHandler) UpdateMember(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid member ID")
		return
	}

	var member models.Member
	if err := decodeJSON(r, &member); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid member data")
		return
	}

	member.ID = id

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.Service.UpdateMember(ctx, &member); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, member)
}

func (h *LibraryHandler) DeleteMember(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid member ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if err := h.Service.DeleteMember(ctx, id); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Member deleted successfully"})
}

// LOAN HANDLERS

type issueRequest struct {
	BookID   int    `json:"bookId"`
	MemberID int    `json:"memberId"`
	DueDate  string `json:"dueDate"`
}

func (h *LibraryHandler) IssueBook(w http.ResponseWriter, r *http.Request) {
	var request issueRequest

	if err := decodeJSON(r, &request); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid loan data")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	dueDate, err := time.Parse("2006-01-02", request.DueDate)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid due date format")
		return
	}
	err = h.Service.IssueBook(ctx, request.BookID, request.MemberID, dueDate)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, map[string]string{"message": "Book issued successfully"})
}

func (h *LibraryHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid loan ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := h.Service.ReturnBook(ctx, id); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Book returned successfully"})
}

func (h *LibraryHandler) GetLoans(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	loans, err := h.Service.GetLoans(ctx)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, loans)
}

func (h *LibraryHandler) GetOverdue(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	loans, err := h.Service.GetOverdueLoans(ctx)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, loans)
}

// BORROWING HISTORY

func (h *LibraryHandler) GetMemberHistory(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid member ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	history, err := h.Service.GetMemberHistory(ctx, id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, history)
}

func (h *LibraryHandler) GetBookHistory(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	history, err := h.Service.GetBookHistory(ctx, id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, history)
}

// DASHBOARD

func (h *LibraryHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var stats struct {
		TotalBooks   int `json:"totalBooks"`
		TotalMembers int `json:"totalMembers"`
		IssuedBooks  int `json:"issuedBooks"`
		OverdueBooks int `json:"overdueBooks"`
	}

	err := h.DB.QueryRow(ctx, `
		SELECT
			(SELECT COALESCE(SUM(total_copies), 0) FROM books),
			(SELECT COUNT(*) FROM members),
			(SELECT COUNT(*) FROM loans WHERE status = 'Issued'),
			(SELECT COUNT(*) FROM loans
			 WHERE status = 'Issued' AND due_date < CURRENT_DATE)
	`).Scan(
		&stats.TotalBooks,
		&stats.TotalMembers,
		&stats.IssuedBooks,
		&stats.OverdueBooks,
	)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Unable to load dashboard statistics")
		return
	}

	respondJSON(w, http.StatusOK, stats)
}
