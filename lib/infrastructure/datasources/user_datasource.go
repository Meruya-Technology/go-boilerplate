package datasources

type UserDatasource interface {
	User() (string, error)
	Login() (int, error)
}
