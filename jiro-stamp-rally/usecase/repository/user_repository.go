package repository

import "github.com/fesInformatics/jiro-stamp-rally/usecase/model"

type UserRepository interface {
	Save(mailAddress string, Password string) error
	ExistsByMailAddress(mailAddress string) (bool, error)
	FindByMailAddress(queryMailAddress string) (*model.User, error)
}
