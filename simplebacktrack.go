package main

import (
	"fmt"
)

var finished bool
var total int

func Backtrack(a []bool, k int, input int) {
	if IsSolution(a, k, input) {
		ProcessSolution(a, k, input)
		return
	}
	k++

	candidates := ConstructCandidates(a, k, input)
	for _, c := range candidates {
		a[k-1] = c
		MakeMove(a, k, input)
		Backtrack(a, k, input)
		UnmakeMove(a, k, input)
		if Finished() {
			return
		}
	}
}

func IsSolution(a []bool, k int, n int) bool {
	return k == n
}

func ConstructCandidates(a []bool, k int, n int) []bool {
	return []bool{true, false}
}

func ProcessSolution(a []bool, k int, n int) {
	fmt.Printf("{")
	for i, v := range a {
		if v {
			fmt.Printf("%d", i)
		}
	}
	fmt.Printf("}\n")
	total++
}

func MakeMove(a []bool, k int, n int) {
}

func UnmakeMove(a []bool, k int, n int) {
}

func Finished() bool {
	return finished
}

func main() {
	a := make([]bool, 10)
	Backtrack(a, 0, 10)
	fmt.Println("total", total)
}
