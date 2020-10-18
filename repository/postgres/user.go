package postgres

import (
	"database/sql"

	"github.com/nayak-rakesh/go-project-setup/domain"
)

type userRepository struct {
	conn *sql.DB
}

// Init ...
func Init(db *sql.DB) domain.UserRepository {
	return &userRepository{
		conn: db,
	}
}

func (repo *userRepository) Create(u *domain.User) (err error) {
	sqlStatement := `
		INSERT INTO users (name, age, email, address)
		VALUES ($1, $2, $3, $4)`
	_, err = repo.conn.Exec(sqlStatement, u.Name, u.Age, u.Email, u.Address)
	if err != nil {
		return
	}
	return nil
}

func (repo *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	sqlStatement := `
		SELECT * FROM users 
		WHERE email = $$`
	row := repo.conn.QueryRow(sqlStatement, email)

	err := row.Scan(&user.Name, &user.Age, &user.Email, &user.Address)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (repo *userRepository) GetAll() ([]domain.User, error) {
	var userList []domain.User
	sqlStatement := `SELECT * FROM users`
	rows, err := repo.conn.Query(sqlStatement)
	if err != nil {
		return userList, err
	}
	user := domain.User{}
	for rows.Next() {
		var name, email, address string
		var age int
		err = rows.Scan(name, age, email, address)
		if err != nil {
			return userList, err
		}
		user.Name = name
		user.Age = age
		user.Email = email
		user.Address = address
		userList = append(userList, user)
	}

	return userList, nil
}
