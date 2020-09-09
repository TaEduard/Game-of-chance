package utils

import (
	"fmt"
	"testing"
)

func TestRandomNoGen01(t *testing.T) {
	ok1 := false
	ok2 := false
	for i := 0; i < 10; i++ {
		number := RandomNoGen(1, 2)
		if number == 1 {
			ok1 = true
		} else if number == 2 {
			ok2 = true
		} else {
			t.Error("The number generated is out of the min-max boundaries.")
		}
	}
	if !ok1 || !ok2 {
		t.Error("The values of the numbers generated didn't vary enough.")
	}
}

func TestRandomNoGen02(t *testing.T) {
	failedAttempts := 0
	for i := 0; i < 10; i++ {
		number1 := RandomNoGen(1, 30)
		number2 := RandomNoGen(1, 30)
		fmt.Println("number1:", number1)
		fmt.Println("number2:", number2)

		if number1 == number2 {
			failedAttempts += 1
		}
	}

	if failedAttempts >= 2 {
		t.Error("The values of the numbers generated didn't vary enough.")
	}
	fmt.Println("failedAttempts:", failedAttempts)
}

func TestRandomNoGen03(t *testing.T) {
	number1 := RandomNoGen(1, 0)
	if number1 != 0 {
		t.Error("The default value is incorrect.")
	}
}
