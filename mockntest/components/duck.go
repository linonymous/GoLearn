package components

type Duck struct {
	Name string
}

func (d Duck) Speak(s string) string {
	return s + d.Name
}

func NewDuck(name string) *Duck {
	return &Duck{Name: name}
}
