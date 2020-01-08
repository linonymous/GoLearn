package components

type Dog struct {
	Name string
}

func (d Dog) Speak(s string) string {
	return s + d.Name
}

func NewDog(name string) *Dog {
	return &Dog{Name: name}
}
