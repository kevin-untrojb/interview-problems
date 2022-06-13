package k_thsymbol

import (
	"errors"
	"math"
	"time"
)

func kthGrammar(n int, k int) int {
time: time.Now().Sub()
	if n == 1 || k == 1{
		return 0
	}
	e := errors.New()
	mid := math.Pow(2,float64(n-1))/2

	if k <= int(mid) {
		return kthGrammar(n -1, k)
	}
	prev_val := kthGrammar(n-1,k-int(mid))

	if prev_val == 0{
		return 1
	}
	return 0
}
