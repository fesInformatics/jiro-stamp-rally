package interactor

import (
	"errors"

	"github.com/fesInformatics/jiro-stamp-rally/usecase/repository"
)

var ErrUserNotFound = errors.New("ユーザーが存在しません")
var ErrNotMatchPassword = errors.New("パスワードが間違っています")

type LoginInteractor struct {
	repository repository.UserRepository
}

func (i *LoginInteractor) Login(mailAddress string, Password string) error {
	user, err := i.repository.FindByMailAddress(mailAddress)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}
	if user.Password != Password {
		return ErrNotMatchPassword
	}
	// TODO トークン生成？APIの認証方法、トークンの保持の仕方などが決まったら実装する
	return nil
}

func NewLoginInteactor(repository repository.UserRepository) LoginInteractor {
	return LoginInteractor{
		repository: repository,
	}
}
