package cat

// CreateAnimal create a new animal given in parameter it type, age and size
func CreateAnimal(animal string, age int, size float32) ICat {
	switch animal {
	case "cat":
		return &cat{age: age, size: size}
	default:
		return nil
	}
}