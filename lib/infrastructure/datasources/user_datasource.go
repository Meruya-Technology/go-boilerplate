package datasources

import mdl "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"

type UserDatasource interface {
	User() (string, error)
	Login() (*mdl.UserModel, error)
	Register() (*mdl.UserModel, error)
	CheckPhone() (bool, error)
	CheckEmail() (bool, error)
}
