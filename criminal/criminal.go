package criminal

type Criminal interface {
	Update(state bool) bool
	GetName() string
}