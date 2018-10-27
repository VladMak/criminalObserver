package main

import (
	"./criminal"
	"fmt"
)

func main() {
	fmt.Println("START")
	var boss = criminal.Boss{Name: "Boss", MainState: false, Subscribers: make(map[string]criminal.Criminal)}

	var porter criminal.Criminal = criminal.Porter{Name: "Porter"}
	var driver criminal.Criminal = criminal.Driver{Name: "Driver"}
	var thief criminal.Criminal = criminal.Thief{Name: "Thief"}

	boss.Subscribe(driver.GetName(), driver)
	boss.Subscribe(porter.GetName(), porter)
	boss.Subscribe(thief.GetName(), thief)

	//porter.Subscribe(boss.GetName(), boss)
	//porter.Subscribe(driver.GetName(), driver)
	//porter.Subscribe(thief.GetName(), thief)

	boss.MainBusinessLogic()
}
