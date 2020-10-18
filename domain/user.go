package domain

// User ...
type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// UserRepository ...
type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetAll() ([]User, error)
}

// UserSerice ...
type UserSerice interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetAll() ([]User, error)
}
