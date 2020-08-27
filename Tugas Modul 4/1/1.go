package main

import "fmt"

func main() {
	var pi, save, bawah float64
	var j, S int
	save = 1
	bawah = 1
	j = 1

	fmt.Print("N Suka pertama : ")
	fmt.Scan(&S)
	for j < S {

		bawah = bawah + 2
		if j%2 != 0 && S != 1 {
			save = save - (1 / bawah)
		}
		if j%2 == 0 && S != 1 {
			save = save + (1 / bawah)
		}
		j = j + 1
	}
	pi = save * 4
	fmt.Print("Hasil Pi : ", pi)
}
