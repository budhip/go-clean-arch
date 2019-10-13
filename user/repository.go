package user

import (
	"github.com/budhip/go-postgre-clean-arch/models"
)

// Repository represent the user's repository contract
type Repository interface {
	FetchUsers() ([]*models.User, error)
}
