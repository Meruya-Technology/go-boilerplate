package mappers

import (
	d "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	m "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type UserMapper struct{}

func (m UserMapper) ToEntity(model m.UserModel) *d.User {
	return &d.User{
		Id:    model.Id,
		Email: model.Email,
	}
}
