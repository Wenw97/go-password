package rand

import (
	"fmt"
	"testing"
)

func TestRandClass_Rand(t *testing.T) {
	class := New([]string{"abcdefghijklmnopqrstuvwxyz", "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "0123456789", "!@#$%^&*()[]{}+-*/_=."})
	str, err := class.Rand(10, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str, len(str))
}
