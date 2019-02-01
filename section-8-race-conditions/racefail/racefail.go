package racefail

import (
	"database/sql"

	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("race: resource could not be located")
)

type User struct {
	ID      int
	Name    string
	Email   string
	Balance int
}

type UserStore struct {
	sql interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		QueryRow(query string, args ...interface{}) *sql.Row
	}
}

func (us *UserStore) Find(id int) (*User, error) {
	var user User
	const query = `SELECT id, name, email, balance FROM users WHERE id=$1;`

	row := us.sql.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Balance)

	switch err {
	case sql.ErrNoRows:
		return nil, ErrNotFound
	case nil:
		return &user, nil
	default:
		return nil, errors.Wrap(err, "race: error querying for user by id")
	}
}

func (us *UserStore) Create(user *User) error {
	const query = `INSERT INTO users (name, email, balance) VALUES ($1, $2, $3) RETURNING id;`
	err := us.sql.QueryRow(query, user.Name, user.Email, user.Balance).Scan(&user.ID)

	if err != nil {
		return errors.Wrap(err, "race: error creating new user")
	}
	return nil
}

func (us *UserStore) Update(user *User) error {
	const query = `UPDATE users SET name=$2, email=$3, balance=$4 WHERE id=$1;`
	_, err := us.sql.Exec(query, user.ID, user.Name, user.Balance)
	if err != nil {
		return errors.Wrap(err, "race: error updating user")
	}
	return nil
}

func (us *UserStore) Delete(id int) error {
	const query = `DELETE FROM users WHERE id = $1;`
	_, err := us.sql.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "race: error deleting user")
	}
	return nil
}

func Spend(us interface {
	Find(int) (*User, error)
	Update(*User) error
}, userID int, amount int) error {
	user, err := us.Find(userID)
	if err != nil {
		return err
	}
	user.Balance -= amount
	return us.Update(user)
}
