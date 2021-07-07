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
	output := RPN("2 3 5 * * 5 +")
	if output != 35 {
		t.Error("Test03 is failed")
	}
}

func TestRPN04(t *testing.T) {
	output := RPN("1 3 / 1 6 / +")
	if output != 1/2 {
		t.Error("Test04 is failed")
	}
}
