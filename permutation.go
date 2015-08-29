package main

import (
	"backtrack"
	"fmt"
)

type PermutationImpl struct {
	finished bool
	total    int
}

func (b *PermutationImpl) IsSolution(a []interface{}, k int, n interface{}) bool {
	return k == n
}

func (b *PermutationImpl) ConstructCandidates(a []interface{}, k int, n interface{}) []interface{} {

	inPerm := make(map[int]bool)

	for i := 0; i < k-1; i++ {
		inPerm[a[i].(int)] = true
	}

	result := []interface{}{}

	for i := 1; i <= n.(int); i++ {
		if !inPerm[i] {
			result = append(result, i)
		}
	}

	return result
}

func (b *PermutationImpl) ProcessSolution(a []interface{}, k int, n interface{}) {
	fmt.Println(a)
	b.total++
}

func (b *PermutationImpl) MakeMove(a []interface{}, k int, n interface{}) {
}

func (b *PermutationImpl) UnmakeMove(a []interface{}, k int, n interface{}) {
}

func (b *PermutationImpl) Finished() bool {
	return b.finished
}

func main() {
	a := make([]interface{}, 4)
	impl := &PermutationImpl{}
	b := backtrack.New(impl)
	b.Backtrack(a, 0, 4)
	fmt.Println("total", impl.total)
}
