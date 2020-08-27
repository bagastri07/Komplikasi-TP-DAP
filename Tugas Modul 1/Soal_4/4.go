package main

import "fmt"

func main() {
	/*
		NIM : 1301194051
		Nama : Bagas Tri Wibowo
		Solusi dari progam "Konvesri Suhu" ini adalah, saya menggunakan perpaduan dari fungsi var dengan
		tipe data float32. Kemudian saya menggunakan fungsi Scanln sebagai perintah input dari suhu celcius
		yang ingin kita konversi, lalu kemudian saya menggunakan fungsi Println, sebagai output dimana
		di dalamnya terdapat perintah operator menggunakan rumus matematika dasar, yakni rumus konversi suhu
	*/
	var celcius float32

	fmt.Println("Masukan Nilai Celcius = ")
	fmt.Scanln(&celcius)
	fmt.Println("Temparatur Celcius=", celcius, "°C")
	fmt.Println("Temparatur Reamur=", celcius*4/5, "°R")
	fmt.Println("Temparatur Fahrenheit=", (celcius*9/5)+32, "°F")
	fmt.Println("Temparatur Kelvin=", (((celcius*9/5)+32)+459.67)*5/9, "K")

	fmt.Println("<=========================>")
	fmt.Println("Nama : Bagas Tri Wibowo")
	fmt.Println("NIM : 1301194051")
}
