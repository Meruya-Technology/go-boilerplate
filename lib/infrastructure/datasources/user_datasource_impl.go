package datasources

type UserDatasourceImpl struct{}

func (impl UserDatasourceImpl) User() string {
	return "Here is your user"
}
