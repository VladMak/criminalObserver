package main

import (
	"./criminal"
	"fmt"
)

func main() {
	fmt.Println("Сделать подмену? (y / n)")
	var boss = criminal.Boss{Name: "Boss", MainState: false, Subscribers: make(map[string]criminal.Criminal)}

	var input string
	fmt.Scanln(&input)

	porter := criminal.Porter{Name: "Porter", Subscribers: make(map[string]criminal.Criminal)}

	if input == "y" {
		porter.PoliceMan = true
	} else {
		porter.PoliceMan = false
	}

	driver := criminal.Driver{Name: "Driver"}
	thief := criminal.Thief{Name: "Thief"}

	boss.Subscribe(driver.GetName(), driver)
	boss.Subscribe(porter.GetName(), porter)
	boss.Subscribe(thief.GetName(), thief)

	porter.Subscribe(boss.GetName(), boss)
	porter.Subscribe(driver.GetName(), driver)
	porter.Subscribe(thief.GetName(), thief)

	boss.MainBusinessLogic()
}
