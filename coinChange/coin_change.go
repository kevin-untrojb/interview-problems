package coinChange

import "math"

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}

	opt := make([]int, amount+1)
	for i := range opt {
		opt[i] = math.MaxInt32
	}
	opt[0]=0
	for am := 1; am <= amount; am++ {
		for _,coin := range coins{
			if coin > am {
				continue
			}
			opt[am] = min(opt[am-coin]+1, opt[am])
		}
	}
	if opt[amount] > amount{
		return -1
	}
	return opt[amount]
}
func min(a int, b int)int{
	if a < b{
		return a
	}
	return b
}