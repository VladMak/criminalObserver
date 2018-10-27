package criminal

import (
	"fmt"
)

type Thief struct {
	Name  string
	State bool
}

func (thief Thief) Update(state bool) bool {
	thief.State = state
	if thief.State {
		fmt.Println(thief.Name, ": Готов, Иду воровать")
		return true
	} else {
		fmt.Println(thief.Name, ": Не готов. Не хочу воровать")
		return true
	}
}

func (thief Thief) GetName() string {
	return thief.Name
}
