package datasources

import "fmt"

type Example interface {
	Something() string
	GetSomething()
}

type ExampleDatasource struct{ Example }

func (datasource ExampleDatasource) GetSomething() {
	fmt.Print(datasource.Something())
}
