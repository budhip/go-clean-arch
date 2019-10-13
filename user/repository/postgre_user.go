package repository

import (
	"database/sql"

	"github.com/sirupsen/logrus"

	"github.com/budhip/go-postgre-clean-arch/models"
	"github.com/budhip/go-postgre-clean-arch/user"
)

type postgreUserRepository struct {
	Conn *sql.DB
}

// NewPostgreUserRepository will create an object that represent the user.Repository interface
func NewPostgreUserRepository(Conn *sql.DB) user.Repository {
	return &postgreUserRepository{Conn}
}

func (pur *postgreUserRepository) fetchUsers(query string) ([]*models.User, error) {
	rows, err := pur.Conn.Query(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errClose := rows.Close()
		if errClose != nil {
			logrus.Error(errClose)
		}
	}()

	result := make([]*models.User, 0)
	for rows.Next() {
		t := new(models.User)
		err = rows.Scan(
			&t.ID,
			&t.FirstName,
			&t.LastName,
			&t.Address,
			&t.DateOfBirth,
			&t.Email,
			&t.AccountConfirmed,
			&t.PhoneNumber,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (pur *postgreUserRepository) FetchUsers() ([]*models.User, error) {
	query := `
		SELECT
			id,
			"firstName",
			"lastName",
			"address",
			"dateOfBirth",
			email,
			"accountConfirmed",
			"phoneNumber",
			"createdAt",
			"updatedAt"
		FROM "Users"`

	res, err := pur.fetchUsers(query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
