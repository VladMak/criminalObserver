package criminal

import (
	"fmt"
)

type Driver struct {
	Name  string
	State bool
}

func (driver Driver) Update(state bool) bool {
	driver.State = state
	if driver.State {
		fmt.Println(driver.Name, ": Готов, Еду")
		return true
	} else {
		fmt.Println(driver.Name, ": Не готов. Машина сломалась")
		return true
	}
}

func (driver Driver) GetName() string {
	return driver.Name
}
