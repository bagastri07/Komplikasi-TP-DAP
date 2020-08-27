/*
	Nama : Bagas Tri Wibowo
	NIM : 1301194051

	Disini saya membuat 4 function, yaitu fungsiF, fungsiG, fungsiH, dan fungsiFoGoH. dan setiap fungsi dibuat sesuai
	permintaan soal. Kemudian semua fungsi akan dipanggil di main code
*/

package main

import "fmt"

func fungsiF(x float32) float32 {
	var fungsiF float32

	fungsiF = x * x
	return fungsiF
}

func fungsiG(x float32) float32 {
	var fungsiG float32

	fungsiG = x - 2
	return fungsiG
}

func fungsiH(x float32) float32 {
	var fungsiH float32

	fungsiH = x + 1
	return fungsiH
}

func fungsiFoGoH(x float32) float32 {
	var fungsiFoGoH float32

	fungsiFoGoH = fungsiF(fungsiG(fungsiH(x)))
	return fungsiFoGoH
}

func main() {
	var x float32

	fmt.Print("Masukan nilai x = ")
	fmt.Scanln(&x)
	fmt.Printf("f(%.2f) = %.2f\n", x, fungsiF(x))
	fmt.Printf("g(%.2f) = %.2f\n", x, fungsiG(x))
	fmt.Printf("h(%.2f) = %.2f\n", x, fungsiH(x))
	fmt.Printf("(fogoh)(%.4f) = %.2f\n", x, fungsiFoGoH(x))

	fmt.Printf("\nNama : Bagas Tri Wibowo")
	fmt.Printf("\nNIM : 1301194051")
}
