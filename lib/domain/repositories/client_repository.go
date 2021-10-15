package repositories

import (
	ent "github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type ClientRepository interface {
	Create() (*ent.Client, error)
}
