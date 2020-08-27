package main

/*
	NIM		: 1301194051
	Nama	: Bagas Tri Wibowo

	Solusi yang saya berikan untuk program ini adalah menggunakan dua variabel bertipe float32
	karena dua inputan akan dinput dengan tipe data float32. Untuk fungsi menggunakan fungsi
	perulangan for "Repeat-Until" dimana progam yang ada di dalamnya akan terus diulang sampai
	nilai inputan bernilai "true"

*/

import "fmt"

func main() {
	var belanja1, belanja2 float32

	for terpenuhi := false; !terpenuhi; {
		fmt.Print("Masukan berat belanjaan di kedua kantong : ")
		fmt.Scan(&belanja1, &belanja2)
		terpenuhi = belanja1 >= 9 || belanja2 >= 9
	}
	fmt.Println("Proses selesai.")

	fmt.Println("\nNama 	: Bagas Tri Wibowo")
	fmt.Println("NIM	: 1301194051")
}
