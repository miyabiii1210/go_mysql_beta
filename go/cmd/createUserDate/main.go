package main

import (
	"context"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/database"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/user"
	"github.com/miyabiii1210/go_mysql_beta/go/pkg/util"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
}

const (
	USER_COUNT       = 100  // 生成するユーザ数
	HOSTNAME_DIGIT   = 8    // emailのホストネームに指定する文字数
	TEL_NUMBER_MAX   = 9999 // 電話番号の第2、第3ブロックの上限値
	TEL_NUMBER_DIGIT = 4    // 電話番号の第2、第3ブロックの桁数
)

func GenerateUserDate(cnt int) error {
	db, err := database.ConnectionToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	u := user.User{}

	// for email address
	multipleRandomStr, err := util.GenerateMultipleRandomString(HOSTNAME_DIGIT, USER_COUNT, util.LowercaseStrAndNumMix)
	if err != nil {
		return err
	}

	// for tel number
	numArr := util.GenerateMultipleRandomNumber(TEL_NUMBER_MAX, USER_COUNT)
	numStr := util.MultipleDigitCompensating(numArr, TEL_NUMBER_DIGIT)

	for i := 0; i < USER_COUNT; i++ {
		u.Name = "test-user-" + strconv.Itoa(i+1)
		u.Email = multipleRandomStr[i] + "@example.co.jp"
		u.TelNumber = "090-" + numStr[i] + "-" + numStr[USER_COUNT-(i+1)]

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
