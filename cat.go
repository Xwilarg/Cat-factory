package cat

// ICat is a representation of a cat that can eat and sleep
type ICat interface {
	GetAge() int
	SetAge(int)
	GetSize() float32
	SetSize(float32)
	Eat()
	Sleep()
}

// cat is an implementation of ICat
type cat struct {
	age  int
	size float32
}
