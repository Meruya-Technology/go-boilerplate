package datasources

type ClientDatasource interface {
	Create() (bool, error)
	Check(secret string, clientId int) (bool, error)
}
