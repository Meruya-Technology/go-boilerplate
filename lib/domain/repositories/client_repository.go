package repositories

import "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"

type ClientRepository interface {
	Create() entities.Client
}
