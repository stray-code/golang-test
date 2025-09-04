package main

import "fmt"

type User struct {
	Name   string
	Height float64
	Weight float64
}

func (u *User) CalculateBMI() (float64, error) {
	if u.Height <= 0 {
		return 0, fmt.Errorf("身長は0より大きい値である必要があります")
	}

	return u.Weight / (u.Height * u.Height), nil
}

func main() {
	user := User{
		Name: "George",
		// Height: 1.75,
		Height: 0,
		Weight: 65.3,
	}

	bmi, error := user.CalculateBMI()
	if error != nil {
		fmt.Printf("BMTの計算に失敗しました：%v\n", error)
	} else {
		fmt.Printf("%sのBMIは%.2fです。\n", user.Name, bmi)
	}

}
