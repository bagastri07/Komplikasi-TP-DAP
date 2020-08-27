package main

import (
	"fmt"
)

const N = 2019

type RecType struct {
	f1 string
	f2 int
	f3 float64
}
type ArrType [N]RecType

func imin(data RecType) int {
	min := 0
	for i := 1; i <= N; i++ {
		if data[min].f2 < data[i].f2 {
			min = i
		}
	}
	return imin
}

func found(data RecType, key string) int {
	for i := 0; i <= N; i++ {
		if data[0].f1 == key {
			return i
		}
	}
}

func main() {
	fmt.Print()
}
