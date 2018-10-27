package criminal

import (
	"fmt"
)

type Porter struct {
	Name        string
	State       bool
	Subscribers map[string]Criminal
	PoliceMan   bool
}

func (porter Porter) Update(state bool) bool {
	porter.State = state
	if porter.PoliceMan {
		porter.MainBusinessLogic()
		return false
	}

	if porter.State {
		fmt.Println(porter.Name, ": Готов, Несу")
		return true
	} else {
		fmt.Println(porter.Name, ": Не готов. Руки болят, не могу ничего нести")
		return true
	}
}

func (porter Porter) GetName() string {
	return porter.Name
}

func (porter *Porter) Test() {
	fmt.Println("TestFunc")
}

func (porter *Porter) Subscribe(name string, criminal Criminal) {
	porter.Subscribers[name] = criminal
}

func (porter *Porter) Unsubscribe(name string) {
	delete(porter.Subscribers, name)
}

func (porter *Porter) NotifySubscribers() {
	fmt.Println(porter.Name, ": Вы все аррестованы")
	for _, v := range porter.Subscribers {
		v.Update(false)
	}
}

func (porter *Porter) MainBusinessLogic() {
	porter.NotifySubscribers()
}
