package user

import (
	"github.com/budhip/go-postgre-clean-arch/models"
)

// Usecase represent the user's usecases
type Usecase interface {
	FetchUsers() ([]*models.User, error)
}
