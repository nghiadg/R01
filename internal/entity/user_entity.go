package entity

type User struct {
	// user identity
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// user information
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
