/*
	Nama : Bagas Tri Wibowo
	NIM : 1301194051

	Saya disini menggunakan perulangan sebanyak 3 perulangan. Dimana diperulangan pertama adalah untuk mencari nilai Pi dari deret
	S ke i, lalu perulangan kedua untuk mencari nilai Pi dari deret S ke (i+1). Dan perulangan ketiga adalah perulangan untuk mengulang
	dua perulangan pertama dan kedua, untuk terus running, sampai nilai selesih nilai pi diantara kedua perulangan kurang dari 0.00001
	lalu terdapat pula fungsi seleksi kondisi di perulangan pertama dan kedua guna menentukan apakah itu Suku ke "GANJIL" atau Suku ke "GENAP"
	dan terakhir terdapat fungsi output, untuk mengoutputkan nilai pi 1, pi 2, dan suku (i) ketika program berhenti
*/

package main

import "fmt"

func main() {
	var pi, pi2, save, bawah float64
	var jenis, Suku, Suku2 int
	save = 1
	bawah = 1
	jenis = 1
	Suku = 1
	Suku2 = Suku + 1
	pi2 = 10
	pi = 1
	for pi2-pi > 0.00001 || pi-pi2 > 0.00001 {
		for jenis < Suku {

			bawah = bawah + 2
			if jenis%2 != 0 && Suku != 1 { //untuk Suku GANJIL
				save = save - (1 / bawah)
			}
			if jenis%2 == 0 && Suku != 1 { //Untuk Suku GENAP
				save = save + (1 / bawah)
			}
			jenis = jenis + 1
		}
		pi = save * 4

		for jenis < Suku2 {

			bawah = bawah + 2
			if jenis%2 != 0 && Suku2 != 1 {
				save = save - (1 / bawah)
			}
			if jenis%2 == 0 && Suku2 != 1 {
				save = save + (1 / bawah)
			}
			jenis = jenis + 1
		}
		pi2 = save * 4
		Suku = Suku + 1
		Suku2 = Suku2 + 1

	}
	fmt.Println("Hasil Pi : ", pi)
	fmt.Println("Hasil Pi : ", pi2)
	fmt.Println("Pada i ke : ", Suku2)
	fmt.Println("<=======================>")
	fmt.Println("Nama : Bagas Tri Wibowo")
	fmt.Println("NIM : 1301194051")
}
