package unittest

import (
	"testing"
)

func TestAdd(t *testing.T) {
	excepted := 5
	actual := Add(2, 3)
	if excepted != actual {
		t.Errorf("excepted：%d, actual:%d", excepted, actual)
	}
}

// gobasic/unittest/add.go
func Add(num1, num2 int) int {
	return num1 + num2
}
