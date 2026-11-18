# Libra – Library Management System

A full-stack Library Management System developed using **Angular, TypeScript, Go (Golang), and PostgreSQL**. The application provides an integrated platform to manage books, library members, borrowing transactions, returns, and library activities through a responsive dashboard.

## Features

### Dashboard
- Displays total books, members, currently issued books, and overdue books.
- Book availability and borrowing status analytics.
- Book collection distribution by category.
- Recent borrowing activity.

### Book Management
- Add, view, update, and delete books.
- Search by title, author, or ISBN.
- Filter by category and sort book records.
- Track total and available copies.
- Prevent deletion of books with borrowing history.

### Member Management
- Register, view, update, and delete members.
- Manage contact details and membership information.
- View member profiles and borrowing history.
- Prevent deletion of members with borrowing history.

### Issue and Return Management
- Issue available books to registered members.
- Record issue dates and due dates.
- Process book returns and update book availability.
- Identify active and overdue loans.
- Maintain borrowing history.

### Additional Features
- Form validation and meaningful error feedback.
- Responsive user interface and navigation.
- Local Storage caching and fallback support.
- PostgreSQL data persistence.
- Reusable Angular components and services.
- REST API architecture and backend business logic.

## Technology Stack

| Layer | Technologies |
|---|---|
| Frontend | Angular, TypeScript, HTML, SCSS |
| Backend | Go (Golang) |
| Database | PostgreSQL |
| API | REST APIs |
| Testing | Go unit tests and API handler tests |
| Version Control | Git and GitHub |

## Project Structure

```text
library-management-system/
│
├── backend/
│   ├── config/
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── services/
│   ├── main.go
│   └── go.mod
│
├── frontend/
│   └── src/
│       └── app/
│           ├── components/
│           ├── pages/
│           ├── services/
│           └── models/
│
├── database/
│   ├── schema.sql
│   └── seed.sql
│
├── screenshots/
├── .gitignore
└── README.md
```

## Installation and Setup

### Prerequisites
- Go
- Node.js and npm
- Angular CLI
- PostgreSQL
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/siyaa10/library-management-system.git
cd library-management-system
```

### 2. Set Up the Database

1. Create a PostgreSQL database named `library_management`.
2. Open pgAdmin and connect to the database.
3. Execute `database/schema.sql` to create the required tables.
4. Execute `database/seed.sql` to insert sample books, members, and borrowing transactions.

**Note:** Run the seed script only once on a fresh database to avoid duplicate records.

### 3. Configure and Run the Backend

Navigate to the backend directory:

```bash
cd backend
```

Create a `.env` file using the variable names provided in `backend/.env.example` and configure your PostgreSQL connection details.

Install dependencies:

```bash
go mod tidy
```

Start the backend:

```bash
go run .
```

Backend API:

`http://localhost:8080`

### 4. Configure and Run the Frontend

Open a separate terminal and navigate to the frontend directory:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Start the Angular application:

```bash
npm start
```

Open the application:

`http://localhost:4200`

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/dashboard` | Dashboard statistics |
| GET / POST | `/api/books` | Retrieve / add books |
| GET / PUT / DELETE | `/api/books/{id}` | View / update / delete books |
| GET / POST | `/api/members` | Retrieve / register members |
| GET / PUT / DELETE | `/api/members/{id}` | View / update / delete members |
| GET | `/api/loans` | Retrieve borrowing records |
| POST | `/api/loans` | Issue a book |
| PUT | `/api/loans/{id}/return` | Return a book |
| GET | `/api/loans/overdue` | Retrieve overdue loans |
| GET | `/api/books/{id}/history` | Book borrowing history |
| GET | `/api/members/{id}/history` | Member borrowing history |

The Books API also supports search, filtering, and sorting.

## Technical Implementation

- **Go Architecture:** Organised using Go modules and packages, with structs, methods, pointers, and interfaces.
- **Business Logic:** Validates inputs, prevents duplicate records, prevents issuing unavailable books, and manages book availability.
- **Concurrency:** Uses goroutines, WaitGroup, and Mutex where appropriate. Database transactions and row-level locking protect borrowing operations.
- **Error Handling:** Provides validation and meaningful API error responses.
- **Data Persistence:** PostgreSQL stores application data, while Local Storage provides frontend caching and fallback support.
- **Testing:** Includes service-level unit tests and API handler tests.

## Testing

Run backend tests from the `backend` directory:

```bash
go test ./...
```

Build the Angular application from the `frontend` directory:

```bash
npm run build
```

## Demo Data

The `database/seed.sql` script provides sample books, members, and borrowing transactions to demonstrate the application's functionality. Seeded records can also be managed through the application, subject to validation and borrowing-history rules.

## Screenshots

Screenshots of the application pages, including the dashboard, book management, member management, issue and return, and detail pages, are available in the `screenshots/` directory.

## Security

- Database credentials are stored in environment variables.
- The `.env` file is excluded from version control.
- Database constraints protect relational data and borrowing records.

---

**Libra – Library Management System**

Built with Angular, Go, and PostgreSQL.