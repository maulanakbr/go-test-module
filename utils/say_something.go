package utils

import (
	"fmt"
	"os"
)

type User struct {
	Name, Address string
}

func SaySomething(name string) string {
	os, _ := os.Hostname()
	user := User{name, os}
	res := fmt.Sprintf("Hi %s, you logged from %s\n", user.Name, user.Address)
	return res
}

/*
To release the original module:
git tag v1.0.0
git push origin v1.0.0
*/
