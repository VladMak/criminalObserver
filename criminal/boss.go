package criminal

import (
	"fmt"
)

type Boss struct {
	Name        string
	Subscribers map[string]Criminal
	MainState   bool
}

func (boss *Boss) Subscribe(name string, criminal Criminal) {
	boss.Subscribers[name] = criminal
}

func (boss *Boss) Unsubsribe(name string) {
	delete(boss.Subscribers, name)
}

func (boss *Boss) NotifySubscribers() {
	fmt.Println(boss.Name, ": Идем воровать")
	for _, v := range boss.Subscribers {
		var bol bool
		if boss.MainState {
			bol = v.Update(true)
		}
		if !bol {
			break
		}

	}
}

func (boss *Boss) MainBusinessLogic() {
	boss.MainState = true
	boss.NotifySubscribers()
}

func (boss Boss) Update(state bool) bool {
	boss.MainState = state

	if boss.MainState {
		fmt.Println(boss.Name, ": Готов, Несу")
		return true
	} else {
		fmt.Println(boss.Name, ": Бежим")
		return false
	}
}

func (boss Boss) GetName() string {
	return boss.Name
}
