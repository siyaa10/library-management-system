package models

import "time"

type Member struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            *string   `json:"email"`
	Phone            string    `json:"phone"`
	Address          string    `json:"address"`
	MembershipType   string    `json:"membershipType"`
	MembershipDate   time.Time `json:"membershipDate"`
	MembershipStatus string    `json:"membershipStatus"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// Check whether the member can borrow books.
func (m *Member) IsActive() bool {
	return m.MembershipStatus == "Active"
}
