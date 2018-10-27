package criminal

type Criminal interface {
	Update(state bool) string
	GetName() string
}