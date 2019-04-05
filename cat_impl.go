package cat

import "fmt"

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
func (c cat) Eat() {
	fmt.Printf("Cat is eating...")
}

// Sleep allow the cat to sleep
func (c cat) Sleep() {
	fmt.Printf("Cat is sleeping...")
}
