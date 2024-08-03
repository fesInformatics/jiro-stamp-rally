package repository

import (
	"fmt"

	database "github.com/fesInformatics/jiro-stamp-rally/infrastructure/gateway/database"

	"github.com/google/uuid"

	"github.com/fesInformatics/jiro-stamp-rally/usecase/model"
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

	defer rows.Close()

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

func (r *UserRepository) FindByMailAddress(queryMailAddress string) (*model.User, error) {
	query := fmt.Sprintf("SELECT id, mail_address, password FROM users WHERE mail_address = \"%s\"", queryMailAddress)

	rows, err := r.client.SelectSQL(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user *model.User
	for rows.Next() {
		var ID, mailAddress, password string
		if err := rows.Scan(&ID, &mailAddress, &password); err != nil {
			return nil, err
		}
		user = &model.User{
			ID:          ID,
			MailAddress: mailAddress,
			Password:    password,
		}
		// 1メールアドレスにつき、1ユーザーしか存在しない想定なのでbreak
		break
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRespository(client database.DbClient) _repository.UserRepository {
	return &UserRepository{client: client}
}
