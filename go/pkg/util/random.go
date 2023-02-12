package util

import (
	"crypto/rand"
	"errors"
	"math"
	math_rand "math/rand"
	"strconv"
	"time"
)

var (
	BaseStr               = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	LowercaseStrOnly      = "abcdefghijklmnopqrstuvwxyz"
	UppercaseStrOnly      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowercaseStrAndNumMix = "abcdefghijklmnopqrstuvwxyz0123456789"
	UppercaseStrAndNumMix = "zABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	ErrGenerateRandomString = errors.New("Failed generate random string")
)

// 0以上, N未満の乱数を生成
func GenerateRandomNumber(digit int) int {
	math_rand.Seed(time.Now().UnixNano())
	randomNum := math_rand.Intn(digit)

	return randomNum
}

// 0以上、N未満の乱数を複数生成
func GenerateMultipleRandomNumber(digit, count int) []int {
	var randomNum int
	var multipleRandomNum []int
	for i := 0; i < count; i++ {
		randomNum = 0
		math_rand.Seed(time.Now().UnixNano() + int64(i))
		randomNum = math_rand.Intn(digit)
		multipleRandomNum = append(multipleRandomNum, randomNum)
	}
	return multipleRandomNum
}

// ランダムな時間遅延
func RandomDelay(digit int) {
	n := GenerateRandomNumber(digit)
	time.Sleep(time.Duration(n) * time.Second)
}

// ランダムな文字列を生成
func GenerateRandomString(digit int, randomType string) (string, error) {
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", ErrGenerateRandomString
	}
	var randomStr string
	for _, v := range b {
		randomStr += string(randomType[int(v)%len(randomType)])
	}

	return randomStr, nil
}

// ランダムな文字列を複数生成
func GenerateMultipleRandomString(digit, count int, randomType string) ([]string, error) {
	var randomStr string
	var multipleRandomStr []string

	for i := 0; i < count; i++ {
		randomStr = ""
		b := make([]byte, digit)
		if _, err := rand.Read(b); err != nil {
			return nil, ErrGenerateRandomString
		}

		for _, v := range b {
			randomStr += string(randomType[int(v)%len(randomType)])
		}

		multipleRandomStr = append(multipleRandomStr, randomStr)
	}

	return multipleRandomStr, nil
}

// 桁数埋め合わせ
func DigitCompensating(num, digit int) string {
	ans := num + int(math.Pow10(digit))
	ret := strconv.Itoa(int(ans))
	return ret[1:]
}

// 複数の数値の桁数埋め合わせ
func MultipleDigitCompensating(num []int, digit int) []string {
	var ret []string
	for _, n := range num {
		ans := n + int(math.Pow10(digit))
		numStr := strconv.Itoa(int(ans))
		ret = append(ret, numStr[1:])
	}

	return ret
}
