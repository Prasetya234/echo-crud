package repository

import (
	"crud_echo/pkg/domain"
	"database/sql"
)

type UserRepsoitory struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepsoitory{
		db: db,
	}
}

func (u UserRepsoitory) CreateUser(req domain.User) error {
	sql := `INSERT INTO users (full_name, email, password, address) values ($1, $2, $3, $4)`
	_, err2 := u.db.Exec(sql, req.FullName, req.Email, req.Password, req.Address)
	if err2 != nil {
		return err2
	}
	return nil
}

func (u UserRepsoitory) UpdateUser(id int, req domain.User) error {
	sql := `UPDATE users SET full_name = $1, email = $2, password = $3, address = $4 WHERE id = $5`
	_, err2 := u.db.Exec(sql, req.FullName, req.Email, req.Password, req.Address, id)
	if err2 != nil {
		return err2
	}
	return nil
}

func (u UserRepsoitory) GetUser() ([]domain.User, error) {
	sql := `SELECT * FROM users`
	rows, err := u.db.Query(sql)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err2 := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Address)
		if err2 != nil {
			return users, err2
		}
		users = append(users, user)
	}
	return users, err
}

func (u UserRepsoitory) GetUserById(id int) (domain.User, error) {
	var user domain.User
	sql := `SELECT * FROM users WHERE id = $1`
	err := u.db.QueryRow(sql, id).Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Address)
	return user, err
}

func (u UserRepsoitory) DeleteUserById(id int) error {
	sql := `DELETE FROM users WHERE id = $1`
	_, err2 := u.db.Exec(sql, id)
	if err2 != nil {
		return err2
	}
	return nil
}

func (u UserRepsoitory) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	sql := `SELECT * FROM users WHERE email = $1`
	err := u.db.QueryRow(sql, email).Scan(&user.Id, &user.FullName, &user.Email, &user.Password, &user.Address)
	return user, err
}
