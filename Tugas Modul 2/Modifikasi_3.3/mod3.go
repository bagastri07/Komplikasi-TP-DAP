package main

/*
	NIM		: 1301194051
	Nama	: Bagas Tri Wibowo

	Solusi dari progam ini ada, saya menggunakan varibel dengan tipe data float32, dimana variabel ini
	akan menjadi input, dan terdapat pula fungsi for (Repeat-Until) dimana
	program akan terus diulang sampai syarat terpenuhi bernilai true.

	Untuk Keterangan motor pak andi oleng atau tidak, saya menggunakan tipe data bool untuk
	mengoutput kan "True" dan "False" dengan operator logika di dalamnya
*/

import "fmt"

func main() {
	var belanja1, belanja2 float32
	var selisih bool

	for terpenuhi := false; !terpenuhi; {
		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scan(&belanja1, &belanja2)
		terpenuhi = belanja1+belanja2 >= 150 || (belanja1 < 0 || belanja2 < 0)
		selisih = (belanja1-belanja2) >= 9 || (belanja2-belanja1) >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng : ", selisih)
	}
	fmt.Println("Proses selesai.")

	fmt.Println("\nNama 	: Bagas Tri Wibowo")
	fmt.Println("NIM	: 1301194051")
}
