package repositories

import (
	"github.com/Meruya-Technology/go-boilerplate/lib/domain/entities"
)

type UserRepository interface {
	GetProfile() entities.User
}
