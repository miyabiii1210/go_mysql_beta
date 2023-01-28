package user

import (
	"context"
	"database/sql"
)

func saveUser(db *sql.DB, ctx context.Context, u *User) (int64, error) {
	query := "INSERT INTO user(uid, name, email, tel_number) VALUES(?, ?, ?, ?);"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, u.Uid, u.Name, u.Email, u.TelNumber)
	if err != nil {
		return 0, err
	}

	ret, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func getUserInfo(db *sql.DB, ctx context.Context, uid uint32) (*User, error) {
	query := "select uid, name, email, tel_number from user where uid = ?;"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, uid)

	r := &User{}
	if err = row.Scan(&r.Uid, &r.Name, &r.Email, &r.TelNumber); err != nil {
		return nil, err
	}

	return r, nil
}
