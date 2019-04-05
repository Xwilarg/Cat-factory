package cat

// ICat is a representation of a cat that can eat and sleep
type ICat interface {
	GetAge() int
	SetAge(int)
	GetSize() float32
	SetSize(float32)
	Eat() string
	Sleep() string
	AddListener(string, chan ICat)
	RemoveListener(string, chan ICat)
	callListeners(string)
}

// cat is an implementation of ICat
type cat struct {
	age       int
	size      float32
	listeners map[string][]chan ICat
}
