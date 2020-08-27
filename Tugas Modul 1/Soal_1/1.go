package main

import "fmt"

func main() {
	/*
		NIM : 1301194051
		Nama : Bagas Tri Wibowo
	*/
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan Input string: ")
	fmt.Scanln(&satu)
	fmt.Println("Masukan Input string: ")
	fmt.Scanln(&dua)
	fmt.Println("Masukan Input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
	fmt.Println("<==========================>")
	fmt.Println("Nama : Bagas Tri Wibowo")
	fmt.Println("NIM : 1301194051")
	fmt.Println("<==========================>")
}
