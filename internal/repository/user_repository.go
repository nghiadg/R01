package repository

import "r01/internal/entity"

type IUserRepository interface {
	FindByEmail(email string) (entity.User, error)
}
