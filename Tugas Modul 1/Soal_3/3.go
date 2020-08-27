package main

import "fmt"

func main() {
	/*
		NIM : 1301194051
		Nama : BAGAS TRI WIBOWO
		Solusi dari progam ini adalah saya memadukan variabel bertipe float64 dengan perintah operator
		penghitungan biasa. Lalu saya mjuga menggunakan perintah Scanln dan Println untuk input dan output
		program ini. Dimana di dalam input terdapat operator perhitungan yang menggunakan rumus dari
		volume bola dan luas kulit bola
	*/

	var jarijari float64
	var pi = 3.1415926535

	fmt.Println("Progam Menghitung Volume dan Luas Kulit Bola")
	fmt.Print("Jari Jari = ")
	fmt.Scanln(&jarijari)
	fmt.Println("Bola dengan jari jari =", jarijari, "memiliki Volume", pi*jarijari*jarijari*jarijari*4/3, "dan memiliki luas kulit", pi*jarijari*jarijari*4)
	fmt.Println(" ")
	fmt.Println("<=======================>")
	fmt.Println("Nama : Bagas Tri Wibowo")
	fmt.Println("NIM : 1301194051")
}
