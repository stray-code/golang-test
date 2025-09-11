package main

import (
	"fmt"

	"github.com/codersgyan/poscast/auth/auth"
	"github.com/codersgyan/poscast/auth/user"
)

func main() {
	auth.LoginWithCredentials("bob", "1234")

	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@example.com",
		Name:  "John Doe",
	}

	fmt.Println(user)

}
