package repository

type UserRepository interface {
	Save(mailAddress string, Password string) error
	ExistsByMailAddress(mailAddress string) (bool, error)
}
