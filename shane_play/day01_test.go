package main

import "testing"

func TestFirst(t *testing.T) {
	First()
}

func TestSecond(t *testing.T) {
	Second()
}
func TestThird(t *testing.T) {
	Third()
}
func TestFizzBuzz(t *testing.T) {
	fizzBuzz(100)
}

func TestPrintA(t *testing.T) {
	PrintA(10)
}

func TestCountStr(t *testing.T) {
	CountStr()
}

func TestReverseStr(t *testing.T) {
	ReverseStr("foobar")
}

func TestAverageSlice(t *testing.T) {
	AverageSlice()
}
