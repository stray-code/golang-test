package main

import "fmt"

type User struct {
	Name   string
	Height float64
	Weight float64
}

func (u *User) CalculateBMI() float64 {
	return u.Weight / (u.Height * u.Height)
}

func main() {
	user := User{
		Name:   "George",
		Height: 1.75,
		Weight: 65.3,
	}

	bmi := user.CalculateBMI()

	fmt.Printf("%sのBMIは%.2fです。\n", user.Name, bmi)
}
