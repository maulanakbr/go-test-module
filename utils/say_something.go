package utils

import (
	"fmt"
	"os"
)

func SaySomething() string {
	os, _ := os.Hostname()
	res := fmt.Sprintf("Hi %s", os)
	return res
}
