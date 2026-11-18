
-- ==========================================
-- LIBRARY MANAGEMENT SYSTEM DATABASE
-- PostgreSQL 18
-- ==========================================

-- 1. BOOKS TABLE
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    author VARCHAR(150) NOT NULL,
    category VARCHAR(100) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    total_copies INTEGER NOT NULL DEFAULT 1
        CHECK (total_copies >= 1),
    available_copies INTEGER NOT NULL DEFAULT 1
        CHECK (available_copies >= 0),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT valid_book_availability
        CHECK (available_copies <= total_copies)
);

-- 2. MEMBERS TABLE
CREATE TABLE IF NOT EXISTS members (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    email VARCHAR(150) UNIQUE,
    phone VARCHAR(20) NOT NULL,
    address TEXT,
    membership_type VARCHAR(50) NOT NULL DEFAULT 'Regular',
    membership_date DATE NOT NULL DEFAULT CURRENT_DATE,
    membership_status VARCHAR(20) NOT NULL DEFAULT 'Active'
        CHECK (membership_status IN ('Active', 'Inactive')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. LOANS TABLE
CREATE TABLE IF NOT EXISTS loans (
    id SERIAL PRIMARY KEY,
    book_id INTEGER NOT NULL REFERENCES books(id),
    member_id INTEGER NOT NULL REFERENCES members(id),
    issue_date DATE NOT NULL DEFAULT CURRENT_DATE,
    due_date DATE NOT NULL,
    returned_at TIMESTAMP,
    status VARCHAR(20) NOT NULL DEFAULT 'Issued'
        CHECK (status IN ('Issued', 'Returned')),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT valid_loan_dates
        CHECK (due_date >= issue_date),

    CONSTRAINT valid_return_status
        CHECK (
            (status = 'Issued' AND returned_at IS NULL)
            OR
            (status = 'Returned' AND returned_at IS NOT NULL)
        )
);

-- 4. PERFORMANCE INDEXES
CREATE INDEX IF NOT EXISTS idx_books_title
    ON books(title);

CREATE INDEX IF NOT EXISTS idx_books_category
    ON books(category);

CREATE INDEX IF NOT EXISTS idx_members_name
    ON members(name);

CREATE INDEX IF NOT EXISTS idx_loans_book
    ON loans(book_id);

CREATE INDEX IF NOT EXISTS idx_loans_member
    ON loans(member_id);

CREATE INDEX IF NOT EXISTS idx_loans_status_due
    ON loans(status, due_date);

-- ==========================================
-- DATABASE SCHEMA CREATED SUCCESSFULLY
-- ==========================================
