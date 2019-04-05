package cat

// GetAge return the cat age
func (c cat) GetAge() int {
	return c.age
}

// SetAge set the cat age
func (c cat) SetAge(age int) {
	c.age = age
}

// GetSize return the cat size
func (c cat) GetSize() float32 {
	return c.size
}

// SetSize return the cat size
func (c cat) SetSize(size float32) {
	c.size = size
}

// Eat allow the cat to eat
func (c cat) Eat() string {
	return "Cat is eating..."
}

// Sleep allow the cat to sleep
func (c cat) Sleep() string {
	return "Cat is sleeping..."
}
