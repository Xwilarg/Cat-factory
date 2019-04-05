package cat

var currFactory animalFactory

type animalFactory struct {
	createAnimalFactoryPtr func(string, int, float32) ICat
}

// CreateAnimal create a new animal given in parameter its type, age and size
func (f *animalFactory) CreateAnimal(animal string, age int, size float32) ICat {
	return f.createAnimalFactoryPtr(animal, age, size)
}

// GetAnimalFactory return the animal factory
func GetAnimalFactory() *animalFactory {
	return &currFactory
}

// UnsetFactory set the default factory as the current one
func (f *animalFactory) UnsetFactory() {
	f.createAnimalFactoryPtr = createAnimalFactoryDefault
}

// SetFactory set a new factory as the current one
func (f *animalFactory) SetFactory(newFactory func(string, int, float32) ICat) {
	f.createAnimalFactoryPtr = newFactory
}

func init() {
	currFactory.createAnimalFactoryPtr = createAnimalFactoryDefault
}

func createAnimalFactoryDefault(animal string, age int, size float32) ICat {
	switch animal {
	case "cat":
		if age <= 0 || size <= 0 {
			return nil
		}
		return &cat{age: age, size: size, listeners: make(map[string][]chan ICat)}
	default:
		return nil
	}
}
