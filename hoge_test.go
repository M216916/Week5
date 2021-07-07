package main

import "testing"

func TestRPN01(t *testing.T) {
	output := RPN("1 1 +")
	if output != 2 {
		t.Error("Test01 is failed")
	}
}

func TestRPN02(t *testing.T) {
	output := RPN("3 1 -")
	if output != 2 {
		t.Error("Test02 is failed")
	}
}

func TestRPN03(t *testing.T) {
	output := RPN("3 6 *")
	if output != 18 {
		t.Error("Test02 is failed")
	}
}
