package mappers

import (
	d "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	m "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"
)

type ClientMapper struct{}

func (m ClientMapper) ToEntity(model m.CLientModel) *d.Client {
	return &d.Client{
		Id:     model.Id,
		Secret: model.Secret,
	}
}
