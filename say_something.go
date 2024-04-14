package go_test_module

import (
	"fmt"
	"os"
)

func SaySomething() string {
	os, _ := os.Hostname()
	res := fmt.Sprintf("Hi %s", os)
	return res
}

/*
git tag v1.0.0
git push origin v1.0.0
*/
