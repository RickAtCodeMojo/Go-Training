package main

import (
	"fmt"
	"math"
)

type vec3 [3][3]float64

func main() {
	//Arrays are only one dimensions
	//You can compose an array of arrays
	var matrix [4][4]bool
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			matrix[i][j] = ((i+j)%2 == 0)
		}
	}
	var vector vec3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			vector[i][j] = float64(i) * float64(j) * math.Pi
		}
	}

	fmt.Println("Matrix: ", matrix)
	fmt.Println("Vector:", vector)

}
