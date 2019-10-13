package models

import "time"

// User represent the user model
type User struct {
	ID               string    `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Address          string    `json:"address"`
	DateOfBirth      string    `json:"dateOfBirth"`
	Email            string    `json:"email"`
	AccountConfirmed bool      `json:"accountConfirmed"`
	PhoneNumber      string    `json:"phoneNumber"`
	ProfilePicture   string    `json:"profilePicture"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
