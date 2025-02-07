package persistence

import (
	"r01/internal/entity"
)

type userPostgresRepository struct {
}

func NewUserPostgresRepository() *userPostgresRepository {
	return &userPostgresRepository{}
}

// implement IUserRepository
func (upr userPostgresRepository) FindByEmail(email string) (entity.User, error) {
	// query to database
	return entity.User{
		ID:       1,
		Email:    "nghiadt.dev@gmail.com",
		Password: "$2a$10$3wNidBE2w.g6BWFzmmi9A.ZW/3dumSVFaDvgomx36Niuu.VUwgf3O", // abcd@1234
	}, nil
}
