//package knowledgecloud

package main

import (
	"fmt"
)

type KnowledgeCloud struct {
	Real bool
	Head DataHead
	Body DataBody
}

func main() {
	A := []int{1, 2, 3, 4}
	B := []int{10, 20, 30, 40}

	C := []int{-1, -2, -3, -4}
	D := []int{-10, -20, -30, -40}

	E := []int{100, 200, 300, 400}
	F := []int{-100, -200, -300, -400}

	//fmt.Println(DataHead{true, 1, 2, 3, 4})
	//fmt.Println(DataBody{A, B, C, D, E, F})

	h := DataHead{true, 1, 2, 3, 3, 4}
	b := DataBody{A, B, B, C, D, E, F}

	fmt.Println(KnowledgeCloud{true, h, b})
}
