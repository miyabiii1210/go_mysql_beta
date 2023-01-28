package env

import (
	"os"
	"strconv"
)

func GetInt(k string, dv int) int {
	if v := os.Getenv(k); len(v) > 0 {
		if ret, err := strconv.Atoi(v); err == nil {
			return ret
		}
	}

	return dv
}
