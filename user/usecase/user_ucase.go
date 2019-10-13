package usecase

import (
	"github.com/budhip/go-postgre-clean-arch/models"
	"github.com/budhip/go-postgre-clean-arch/user"
)

type userUsecase struct {
	userRepo user.Repository
}

// NewUserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(u user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) FetchUsers() ([]*models.User, error) {
	listUser, err := u.userRepo.FetchUsers()
	if err != nil {
		return listUser, err
	}

	return listUser, nil
}
