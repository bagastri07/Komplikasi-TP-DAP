package main

import "fmt"

func main() {
	var pi, pi2, save, save2, bawah, bawah2 float64
	var j, j2, S, S2 int
	save = 1
	save2 = 1
	bawah = 1
	bawah2 = 1
	j = 1
	j2 = 1
	S = 20000
	S2 = S + 1
	pi2 = 10
	pi = 1
	for pi2-pi > 0.00001 || pi-pi2 > 0.00001 {
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

		for j2 < S2 {

			bawah2 = bawah2 + 2
			if j2%2 != 0 && S2 != 1 {
				save2 = save2 - (1 / bawah2)
			}
			if j2%2 == 0 && S2 != 1 {
				save2 = save2 + (1 / bawah2)
			}
			j2 = j2 + 1
		}
		pi2 = save2 * 4
		S = S + 1
		S2 = S2 + 1
	}
	fmt.Println("Hasil Pi : ", pi)
	fmt.Println("Hasil Pi : ", pi2)
	fmt.Println("S : ", S+1)
}
