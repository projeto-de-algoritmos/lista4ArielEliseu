package main

import (
	"./knowledgecloud"
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 4}

	B := []int{10, 20, 30, 40}
	C := []int{-10, -20, -30, -40}

	D := []int{-1, -2, -3, -4}
	E := []int{-10, -20, -30, -40}

	F := []int{100, 200, 300, 400}

	//fmt.Println(knowledgecloud.DataHead{true, 1, 2, 3, 3, 4})
	//fmt.Println(knowledgecloud.DataBody{A, B, C, D, E, F, G})

	h := knowledgecloud.DataHead{true, 1, 2, 3, 3, 4}
	b := knowledgecloud.DataBody{A, B, B, C, D, E, F}

	fmt.Println(knowledgecloud.Cloud{true, h, b})
}
