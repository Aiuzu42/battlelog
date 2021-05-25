package roll

import (
	"math/rand"
	"time"
)

type Result struct {
	Results []int `json:"results"`
	Sorted  []int `json:"sorted"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Roll(n int, s int) Result {
	rr := Result{Sorted: make([]int, s), Results: make([]int, n)}
	for i := 0; i < n; i++ {
		d := 1 + rand.Intn(s)
		rr.Results[i] = d
		rr.Sorted[d-1] = rr.Sorted[d-1] + 1
	}
	return rr
}
