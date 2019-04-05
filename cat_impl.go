package cat

// GetAge return the cat age
func (c *cat) GetAge() int {
	return c.age
}

// SetAge set the cat age
func (c *cat) SetAge(age int) {
	c.age = age
	c.callListeners("age")
}

// GetSize return the cat size
func (c *cat) GetSize() float32 {
	return c.size
}

// SetSize return the cat size
func (c *cat) SetSize(size float32) {
	c.size = size
	c.callListeners("size")
}

// Eat allow the cat to eat
func (c *cat) Eat() string {
	defer c.callListeners("eat")
	return "Cat is eating..."
}

// Sleep allow the cat to sleep
func (c *cat) Sleep() string {
	defer c.callListeners("sleep")
	return "Cat is sleeping..."
}

// AddListener add a listener on a function
func (c *cat) AddListener(name string, callback chan ICat) {
	if _, ok := c.listeners[name]; ok {
		c.listeners[name] = append(c.listeners[name], callback)
	} else {
		c.listeners[name] = []chan ICat{callback}
	}
}

// RemoveListener remove a listener on a function
func (c *cat) RemoveListener(name string, callback chan ICat) {
	if elems, ok := c.listeners[name]; ok {
		for i, e := range elems {
			if e == callback {
				c.listeners[name] = append(c.listeners[name][:i], c.listeners[name][i+1:]...)
			}
		}
	}
}

func (c *cat) callListeners(name string) {
	if elems, ok := c.listeners[name]; ok {
		for _, e := range elems {
			e <- c
		}
	}
}
