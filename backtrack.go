package backtrack

type Backtracker struct {
	impl BacktrackerImpl
}

type BacktrackerImpl interface {
	IsSolution(a []interface{}, k int, input interface{}) bool

	ConstructCandidates(a []interface{}, k int, n interface{}) []interface{}

	ProcessSolution(a []interface{}, k int, n interface{})

	MakeMove(a []interface{}, k int, n interface{})

	UnmakeMove(a []interface{}, k int, n interface{})

	Finished() bool
}

func New(i BacktrackerImpl) *Backtracker {
	return &Backtracker{impl: i}
}

func (b Backtracker) Backtrack(a []interface{}, k int, input interface{}) {
	if b.impl.IsSolution(a, k, input) {
		b.impl.ProcessSolution(a, k, input)
		return
	}
	k++

	candidates := b.impl.ConstructCandidates(a, k, input)
	for _, c := range candidates {
		a[k-1] = c
		b.impl.MakeMove(a, k, input)
		b.Backtrack(a, k, input)
		b.impl.UnmakeMove(a, k, input)
		if b.impl.Finished() {
			return
		}
	}
}
