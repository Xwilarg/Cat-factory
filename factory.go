package cat

// CreateAnimal create a new animal given in parameter it type, age and size
func CreateAnimal(animal string, age int, size float32) ICat {
	switch animal {
	case "cat":
		if age <= 0 || size <= 0 {
			return nil
		}
		return &cat{age: age, size: size, listeners: make(map[string][]*func())}
	default:
		return nil
	}
}
