package coinChange

import (
	"math"

	"github.com/kevin-untrojb/interview-problems/utils"
)

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	opt := make([]int, amount+1)
	for i := range opt {
		opt[i] = math.MaxInt32
	}
	opt[0] = 0
	for am := 1; am <= amount; am++ {
		for _, coin := range coins {
			if coin > am {
				continue
			}
			opt[am] = utils.Min(opt[am-coin]+1, opt[am])
		}
	}
	if opt[amount] > amount {
		return -1
	}
	return opt[amount]
}
