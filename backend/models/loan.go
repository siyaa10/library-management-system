package models

import "time"

type Loan struct {
	ID         int        `json:"id"`
	BookID     int        `json:"bookId"`
	MemberID   int        `json:"memberId"`
	IssueDate  time.Time  `json:"issueDate"`
	DueDate    time.Time  `json:"dueDate"`
	ReturnedAt *time.Time `json:"returnedAt"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
}

// Check whether a loan is overdue.
func (l *Loan) IsOverdue() bool {
	today := time.Now().Format("2006-01-02")
	dueDate := l.DueDate.Format("2006-01-02")

	return l.Status == "Issued" && today > dueDate
}
