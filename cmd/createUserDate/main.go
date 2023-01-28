package main

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/database"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/user"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}

const (
	USER_COUNT = 100
)

func GenerateUserDate(cnt int) error {
	db, err := database.ConnectionToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	u := user.User{
		Name:      "test-user",
		Email:     "test-mail@example.co.jp",
		TelNumber: "050-1234-5678",
	}

	for i := 0; i < USER_COUNT; i++ {
		lastUid, err := user.SaveUser(db, context.TODO(), &u)
		if err != nil {
			return err
		}
		fmt.Println("save user completed.", lastUid)
	}

	return nil
}

func main() {
	if err := GenerateUserDate(USER_COUNT); err != nil {
		panic(err)
	}
	return
}
