package main

import (
	"fmt"
	"math"
)

//Point adalah tipe data yang berisi x dan y
type Point struct {
	x, y float64
}

var points [10000]Point
var numpoints int

func jarak(p1 Point, p2 Point) float64 {
	var val float64
	val = (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)

	return math.Sqrt(val)
}

func bacaTitik() {
	numpoints = 0
	fmt.Scan(&points[numpoints].x, &points[numpoints].y)
	for (points[numpoints].x != 0.0 || points[numpoints].y != 0.0) && numpoints < 10000 {
		numpoints = numpoints + 1
		fmt.Scan(&points[numpoints].x, &points[numpoints].y)
	}
}

func ambilTitikTerdekat(p1, p2 *Point) {
	var Minimum float64

	Minimum = jarak(points[0], points[1])
	*p1 = points[0]
	*p2 = points[1]
	for i := 0; i < numpoints; i++ {
		for j := i + 1; j < numpoints; j++ {
			if jarak(points[i], points[j]) < Minimum {
				Minimum = jarak(points[i], points[j])
				*p1 = points[i]
				*p2 = points[j]
			}

		}
	}
}

func main() {
	var Ttk1, Ttk2 Point

	bacaTitik()
	ambilTitikTerdekat(&Ttk1, &Ttk2)
	fmt.Printf("\nTitik terdekat adalah (%.1f,%.1f) dan (%.1f,%.1f) dengan jarak %.1f\n", Ttk1.x, Ttk1.y, Ttk2.x, Ttk2.y, jarak(Ttk1, Ttk2))
	fmt.Printf("\nNama : Bagas Tri Wibowo \nNIM : 1301194051\n")
}
