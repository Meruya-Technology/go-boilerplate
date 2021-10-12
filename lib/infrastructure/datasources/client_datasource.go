package datasources

import m "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/models"

type ClientDatasource interface {
	Create() (m.CLientModel, error)
	Check(secret string, clientId int) (bool, error)
}
