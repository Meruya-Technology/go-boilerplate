package datasources

type ExampleDatasourceImpl struct{ ExampleDatasource }

func (impl ExampleDatasourceImpl) Something() string {
	return "here is your something"
}
