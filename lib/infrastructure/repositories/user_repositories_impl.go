package repositories_impl

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
	ds "github.com/Meruya-Technology/go-boilerplate/lib/infrastructure/datasources"
)

type UserRepositoriesImpl struct{}

func (repo UserRepositoriesImpl) GetProfile() entities.User {
	var newDs ds.UserDatasource
	newDs = new(ds.UserDatasourceImpl)
	newDs.User()
	return entities.User{}
}
