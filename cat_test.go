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
