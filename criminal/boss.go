package criminal

import(
	"fmt"
)

type Boss struct {
	Name string
	Subscribers map[string]Criminal
	MainState bool
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
		fmt.Println(v.GetName(), ":",v.Update(true))
	}
}

func (boss *Boss) MainBusinessLogic() {
	boss.MainState = true
	boss.NotifySubscribers()
}

func (boss Boss) Update(state bool) string{
	boss.MainState = state

	if boss.MainState {
		return "Готов, Несу"
	} else {
		return "Бежим"
	}
}

func (boss Boss) GetName() string{
	return boss.Name
}
