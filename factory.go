package cat

import "sync"

var currFactory animalFactory
var mutex sync.Mutex

type animalFactory struct {
	createAnimalFactoryPtr func(string, int, float32) ICat
}

// CreateAnimal create a new animal given in parameter its type, age and size
func (f *animalFactory) CreateAnimal(animal string, age int, size float32) ICat {
	mutex.Lock()
	defer mutex.Unlock()
	return f.createAnimalFactoryPtr(animal, age, size)
}

// GetAnimalFactory return the animal factory
func GetAnimalFactory() *animalFactory {
	mutex.Lock()
	defer mutex.Unlock()
	return &currFactory
}

// UnsetFactory set the default factory as the current one
func (f *animalFactory) UnsetFactory() {
	mutex.Lock()
	f.createAnimalFactoryPtr = createAnimalFactoryDefault
	mutex.Unlock()
}

// SetFactory set a new factory as the current one
func (f *animalFactory) SetFactory(newFactory func(string, int, float32) ICat) {
	mutex.Lock()
	f.createAnimalFactoryPtr = newFactory
	mutex.Unlock()
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
