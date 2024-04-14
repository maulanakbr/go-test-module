package utils

import (
	"fmt"
	"os"
)

type User struct {
	Name, Address string
}

func SaySomething(user User) string {
	os, _ := os.Hostname()
	newUser := User{user.Name, user.Address}
	res := fmt.Sprintf("Hi %s, you logged from %s\nYour host is %s", newUser.Name, newUser.Address, os)
	return res
}

/*
To release/update the original module:
git tag v1.0.0
git push origin v1.0.0

If you want to make major changes:
Change the module's name
for ex: module github.com/maulanakbr/go-test-module/v2

Notes:
Always update the tag when changes are made
for ex: v1.1.1
*/
