package main

import (
	"backtrack"
	"fmt"
)

type SubsetImpl struct {
	finished bool
	total    int
}

func (b *SubsetImpl) IsSolution(a []interface{}, k int, n interface{}) bool {
	return k == n
}

func (b *SubsetImpl) ConstructCandidates(a []interface{}, k int, n interface{}) []interface{} {
	result := make([]interface{}, 2)
	result[0] = true
	result[1] = false
	return result
}

func (b *SubsetImpl) ProcessSolution(a []interface{}, k int, n interface{}) {
	fmt.Printf("{")
	for i, v := range a {
		if v.(bool) {
			fmt.Printf("%d", i)
		}
	}
	fmt.Printf("}\n")
	b.total++
}

func (b *SubsetImpl) MakeMove(a []interface{}, k int, n interface{}) {
}

func (b *SubsetImpl) UnmakeMove(a []interface{}, k int, n interface{}) {
}

func (b *SubsetImpl) Finished() bool {
	return b.finished
}

func main() {
	a := make([]interface{}, 10)
	impl := &SubsetImpl{}
	b := backtrack.New(impl)
	b.Backtrack(a, 0, 10)
	fmt.Println("total", impl.total)
}
