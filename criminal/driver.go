package criminal

type Driver struct {
	Name string
	State bool
}

func (driver Driver) Update(state bool) string {
	driver.State = state
	if driver.State {
		return "Готов, Еду"
	} else {
		return "Не готов. Машина сломалась"
	}
}

func (driver Driver) GetName() string{
	return driver.Name
}