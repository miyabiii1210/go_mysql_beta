package user

import (
	"context"
	"database/sql"
	"errors"
)

type User struct {
	Uid       uint32
	Name      string
	Email     string
	TelNumber string
}

var (
	ErrUserNotFound = errors.New("user not found.")
)

func SaveUser(db *sql.DB, ctx context.Context, u *User) (int64, error) {
	return saveUser(db, ctx, u)
}

func GetUserInfo(db *sql.DB, ctx context.Context, uid uint32) (*User, error) {
	u, err := getUserInfo(db, ctx, uid)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}

func GetAllUsersInfo(db *sql.DB, ctx context.Context) ([]*User, error) {
	return getAllUsersInfo(db, ctx)
}

func DeleteUser(db *sql.DB, ctx context.Context, uid uint32) error {
	return deleteUser(db, ctx, uid)
}
