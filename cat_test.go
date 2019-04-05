package cat

import (
	"testing"
)

func checkFloat(t *testing.T, result, expectation float32, currentTest string) {
	if result < expectation-0.1 || result > expectation+0.1 {
		t.Errorf("Incorrect %s, got: %f, expected: %f", currentTest, result, expectation)
	}
}

func checkInt(t *testing.T, result, expectation int, currentTest string) {
	if result != expectation {
		t.Errorf("Incorrect %s, got: %d, expected: %d", currentTest, result, expectation)
	}
}

func checkString(t *testing.T, result, expectation, currentTest string) {
	if result != expectation {
		t.Errorf("Incorrect %s, got: %s, expected: %s", currentTest, result, expectation)
	}
}

func checkNil(t *testing.T, result ICat, currentTest string) {
	if result != nil {
		t.Errorf("Incorrect %s, expected (nil)", currentTest)
	}
}

func checkNotNil(t *testing.T, result ICat, currentTest string) {
	if result == nil {
		t.Errorf("Incorrect %s, expected not (nil)", currentTest)
	}
}

func TestCatFactory(t *testing.T) {
	c := CreateAnimal("cat", 10, 10)
	checkNotNil(t, c, "Factory")
}

func TestCatFactoryError1(t *testing.T) {
	c := CreateAnimal("cat", -10, 10)
	checkNil(t, c, "Factory error")
}

func TestCatFactoryError2(t *testing.T) {
	c := CreateAnimal("cat", 10, -10)
	checkNil(t, c, "Factory error")
}

func TestGetAge(t *testing.T) {
	c := cat{age: 10, size: 10}
	checkInt(t, c.GetAge(), 10, "GetAge")
}

func TestSetAge(t *testing.T) {
	c := cat{age: 10, size: 10}
	c.SetAge(20)
	checkInt(t, c.age, 10, "SetAge")
}

func TestGetSize(t *testing.T) {
	c := cat{age: 10, size: 10}
	checkFloat(t, c.GetSize(), 10, "GetSize")
}

func TestSetSize(t *testing.T) {
	c := cat{age: 10, size: 10}
	c.SetSize(20)
	checkFloat(t, c.size, 10, "SetSize")
}

func TestEat(t *testing.T) {
	c := cat{age: 10, size: 10}
	checkString(t, c.Eat(), "Cat is eating...", "Eat")
}

func TestSleep(t *testing.T) {
	c := cat{age: 10, size: 10}
	checkString(t, c.Sleep(), "Cat is sleeping...", "Sleep")
}
