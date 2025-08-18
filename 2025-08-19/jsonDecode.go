package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	IsAdmin bool     `json:"isAdmin"`
	Emails  []string `json:"emails"`
}

func main() {
	jsonData := `{
		"name": "Takeshi",
		"age": 34,
		"isAdmin": true,
		"emails": [
			"takeshi@example.com",
			"kasumi@example.com"
		]
	}`

	var user User

	err := json.Unmarshal([]byte(jsonData), &user)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("User Name: %s\n", user.Name)
	fmt.Printf("Age: %d\n", user.Age)
	fmt.Printf("IsAdmin: %t\n", user.IsAdmin)
	fmt.Printf("Emails: %v\n", user.Emails)
}
