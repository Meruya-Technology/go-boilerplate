package repositories_impl

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type UserRepositoriesImpl struct{}

func (repo UserRepositoriesImpl) GetProfile() entities.User {
	return entities.User{
		Id:    1,
		Name:  "Jhon",
		Email: "jhondoe@gmail.com",
	}
}
