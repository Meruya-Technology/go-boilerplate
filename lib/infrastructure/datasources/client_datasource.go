package datasources

import m "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"

type ClientDatasource interface {
	Create(name string) (*m.ClientModel, error)
}
