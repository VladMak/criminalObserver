package criminal

type Thief struct {
	Name string
	State bool
}

func (thief Thief) Update(state bool) string {
	thief.State = state
	if thief.State {
		return "Готов, Иду воровать"
	} else {
		return "Не готов. Не хочу воровать"
	}
}

func (thief Thief) GetName() string{
	return thief.Name
}