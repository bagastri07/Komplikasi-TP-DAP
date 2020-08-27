package main

import "fmt"

func main() {
	/*
		NIM : 1301194051
		NAMA : Bagas Tri Wibowo

		Solusi yang saya berikan untuk membuat progam "Menentukan Tahun Kabisat"
		adalah dengan memadukan fungsi dua fungsi var dengan tipe data boolean,
		yanki kriteria 1 yang berisi operator logika dimana nilainya akan bernilai true
		jika input "tahun" habis dibagi 400 dan kriteria 2 yang berisi operator logika,
		yang akan bernilai true, jika input "tahun" habis dibagi 4 tetapi tidak habis dibagi 100.
		Setelah itu terdapat fungsi Scanln yang berfungsi untuk menginput "tahun" yang diingkan.
		Dan diakhir, terdapat fungsi Println yang berfungsi sebagai output, dimana di dalam
		fungsi tersebut terdapat operator logika, yang berisi jika salah satu kriteria atau kedua kriteria
		bernilai true, makan tahun tersebut akan dioutput sebagai Kabisat=true, jika tidak maka
		tahun tersebut akan dioutput sebagai kabisat=false
	*/
	var tahun int

	fmt.Println("Nama : Bagas Tri Wibowo")
	fmt.Println("NIM : 1301194051")
	fmt.Println(" ")
	fmt.Println("Progam Menentukan Tahun Kabisat")
	fmt.Print("tahun = ")
	fmt.Scanln(&tahun)
	fmt.Println("======>Hasil<======")

	var kriteria1 bool
	var kriteria2 bool
	kriteria1 = (tahun%400 == 0)
	kriteria2 = (tahun%4 == 0 && tahun%100 != 0)

	fmt.Println("Tahun =", tahun)
	fmt.Println("Kabisat =", kriteria1 || kriteria2)
}
