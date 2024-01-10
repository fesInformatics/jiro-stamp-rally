package repository

import (
	"fmt"
	database "github.com/fesInformatics/jiro-stamp-rally/infrastructure/gateway/database"

	"github.com/google/uuid"

	_repository "github.com/fesInformatics/jiro-stamp-rally/usecase/repository"
)

type UserRepository struct {
	client database.DbClient
}

func (r *UserRepository) Save(mailAddress string, Password string) error {
	valuesMap := map[string]any{"id": uuid.New().String(), "mail_address": mailAddress, "password": Password}
	err := r.client.InsertSQL("users", valuesMap)
	if err != nil {
		return err
	}
	return err
}

func (r *UserRepository) ExistsByMailAddress(mailAddress string) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM users WHERE mail_address = \"%s\"", mailAddress)

	rows, err := r.client.SelectSQL(query)
	if err != nil {
		return false, err
	}

	exists := false
	for rows.Next() {
		// 1行でも存在すればユーザーが存在するのでScanする必要がない
		exists = true
	}
	if err := rows.Close(); err != nil {
		return false, err
	}

	return exists, nil
}

func NewUserRespository(client database.DbClient) _repository.UserRepository {
	return &UserRepository{client: client}
}
