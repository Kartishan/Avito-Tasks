package data

import (
	"database/sql"
	"errors"
	"fmt"
)

func (u UserModel) Create() {
	query := `
		INSERT INTO usert (user_cash, user_reserved_cash)
		VALUES (0, 0)`

	u.DB.QueryRow(query).Scan()
}

func (u UserModel) CreateId(id int64) {
	query := `
		INSERT INTO usert (user_id, user_cash, user_reserved_cash)
		VALUES ($1, 0, 0)`

	u.DB.QueryRow(query, id).Scan()
}

func (u UserModel) Get(id int64) (*User, error) {
	if id < 1 {
		return nil, errors.New("incorrect id")
	}

	query := `
			SELECT user_id, user_cash, user_reserved_cash
			FROM usert
			WHERE user_id = $1
		`

	var user User
	err := u.DB.QueryRow(query, id).Scan(
		&user.UserId,
		&user.UserCash,
		&user.UserReservedCash,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("xz")
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (u UserModel) Update(user *User) error {
	query := `
		UPDATE usert
		SET user_cash = $2
		WHERE user_id = $1;
`
	args := []interface{}{
		user.UserId,
		user.UserCash,
	}

	return u.DB.QueryRow(query, args...).Scan()
}

func (u UserModel) UpdateFull(user *User) error {
	fmt.Println("UpdateFull")
	query := `
		UPDATE usert
		SET user_cash = $2, user_reserved_cash = $3
		WHERE user_id = $1;
`
	args := []interface{}{
		user.UserId,
		user.UserCash,
		user.UserReservedCash,
	}

	return u.DB.QueryRow(query, args...).Scan()
}
