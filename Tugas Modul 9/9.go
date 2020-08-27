package main

import (
	"fmt"
)

const MAXSIZE = 2019202

type RecType struct {
	count, size int
}
type ArrType [MAXSIZE]RecType

func iSort(tab *ArrType, nsize int) {
	var i, t, j int

	i = 1
	for i < nsize {
		t = tab[i].count
		j = i - 1
		for j >= 0 && tab[j].count > t {
			tab[j+1].count = tab[j].count
			j = j - 1
		}
		tab[j+1].count = t
		i = i + 1
	}
}

func mSort(tab *ArrType, nsize int) {

	for i := 0; i < nsize-1; i++ {
		Min := i
		for j := i + 1; j <= nsize-1; j++ {
			if tab[j].size > tab[Min].size {
				Min = j
			}
		}
		t := tab[Min].size
		tab[Min].size = tab[i].size
		tab[i].size = t
	}

}

func isFound(tab ArrType, n int, v int) bool {
	kr := 0
	kn := n
	found := false
	for kr < kn && !found {
		mid := (kr + kn) / 2
		if tab[mid].count < v {
			kr = mid + 1
		} else if tab[mid].count > v {
			kn = mid
		} else {
			found = true
		}
	}
	return found
}

func main() {
	var data ArrType
	var n, key int

	fmt.Printf("Jumlah data : ")
	fmt.Scan(&n)
	fmt.Printf("Data Count :\n")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].count)
	}
	fmt.Printf("Data size : \n")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i].size)
	}

	iSort(&data, n)
	fmt.Printf("\nList Count sesudah diurutkan\n")
	for i := 0; i < n; i++ {
		fmt.Print(data[i].count, " ")
	}

	mSort(&data, n)
	fmt.Printf("\nList size sesudah terurut\n")
	for i := 0; i < n; i++ {
		fmt.Print(data[i].size, " ")
	}

	fmt.Printf("\nNilai key data count yang dicari : \n")
	fmt.Scan(&key)

	status := isFound(data, n, key)
	fmt.Printf("\nApakah nilai tersebut ada ?\n%v\n", status)
	fmt.Printf("\nNama :Bagas Tri Wibowo\nNIM :1301194051")
}
