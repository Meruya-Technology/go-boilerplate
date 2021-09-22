package repositories

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type UserRepositories interface {
	GetProfile() entities.User
}
