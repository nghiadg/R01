package persistence

import (
	"r01/internal/entity"
)

type UserPostgresRepository struct {
}

func NewUserPostgresRepository() *UserPostgresRepository {
	return &UserPostgresRepository{}
}

// implement IUserRepository
func (upr UserPostgresRepository) FindByEmail(email string) (entity.User, error) {
	// query to database
	return entity.User{
		ID:        1,
		Email:     "nghiadt.dev@gmail.com",
		Password:  "$2a$10$3wNidBE2w.g6BWFzmmi9A.ZW/3dumSVFaDvgomx36Niuu.VUwgf3O", // abcd@1234
		FirstName: "Nghia",
		LastName:  "Duong",
	}, nil
}
