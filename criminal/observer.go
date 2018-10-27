package criminal

type Observer interface {
	Subscribe(name string, criminal Criminal)
	Unsubsribe(name string)
	NotifySubscribers()
	MainBusinessLogic()
}