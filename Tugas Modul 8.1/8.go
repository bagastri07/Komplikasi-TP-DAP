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

func rmax(data ArrType) float64 {
	rmax := data[0].f3
	for i := 1; i < N; i++ {
		if rmax < data[i].f3 {
			rmax = data[i].f3
		}
	}
	return rmax
}

func imin(data ArrType) int {
	min := 0
	for i := 1; i < N; i++ {
		if data[min].f2 > data[i].f2 {
			min = i
		}
	}
	return min
}

func found(data ArrType, key string) bool {
	for i := 0; i < N; i++ {
		if data[i].f1 == key {
			return true
		}
	}
	return false
}

func pos(data ArrType, key string) int {
	for i := 0; i < N; i++ {
		if data[i].f1 == key {
			return i
		}
	}
	return -1
}

func main() {
	var X ArrType

	i := 0
	fmt.Scan(&X[0].f1, &X[0].f2, &X[0].f3)
	for X[i].f1 != "0" || X[i].f2 != 0 || X[i].f3 != 0 {
		i++
		fmt.Scan(&X[i].f1, &X[i].f2, &X[i].f3)
	}
	key := "ada"
	p := rmax(X)
	q := imin(X)
	r := found(X, key)
	s := pos(X, key)

	fmt.Printf("\nNilai Real Terbesar : %v\nIndex Integer Terkecil : %v\nApakah ada Key? %v\nIndex Key : %v", p, q, r, s)
	fmt.Printf("\n\nNama : Bagas Tri Wibowo\nNIM : 1301194051\nKelas : IF-43-04\n")

}
