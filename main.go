package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func MakeProfile(Name string) Profile {
	var profile Profile
	profile.Name = Name

	if profile.Name == "Goku" {
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
	}

	if profile.Name == "Naruto" {
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
	}

	if profile.Name == "Sasuke" {
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
	}

	return profile
}

func PowerUp(p Profile, i int) Profile {
	var PowerUp Profile

	PowerUp.Name = p.Name
	PowerUp.Health = p.Health * i
	PowerUp.Power = p.Power * i
	PowerUp.Exp = p.Exp * i

	return PowerUp
}

func main() {
	profile := MakeProfile("Sasuke")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power : ", profile.Power)
	fmt.Println("Exp : ", profile.Exp)
	fmt.Println("<===== POWER UP CHARACTER ======>")
	power := PowerUp(profile, 2)
	fmt.Println("Name : ", power.Name)
	fmt.Println("Health : ", power.Health)
	fmt.Println("Power : ", power.Power)
	fmt.Println("Exp : ", power.Exp)

}
