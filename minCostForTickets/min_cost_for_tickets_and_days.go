package minCostForTickets

import (
	"github.com/kevin-untrojb/interview-problems/utils"
	"math"
)

func MincostTicketsWithDaysDp(days []int, costs []int, cantDays []int) int {
	opt := make([]int, len(days))
	costForDay := make([]int, len(cantDays))
	for i := range opt {
		opt[i] = math.MaxInt32
	}
	prevDayCost := 0
	for i, day := range days {
		for i := range costForDay {
			costForDay[i] = prevDayCost + costs[i]

			for j := i + 1; j < len(days); j++ {
				if days[j] > day+cantDays[i]-1 {
					break
				}
				opt[j] = utils.Min(opt[j], costForDay[i])
			}
		}

		prevDayCost = minInTheArray(costForDay)
		opt[i], prevDayCost = utils.Min(opt[i], prevDayCost), utils.Min(opt[i], prevDayCost)
	}
	return opt[len(days)-1]
}

func min3(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

func minInTheArray(arr []int) int {
	if len(arr) < 1 {
		return 0
	}
	max := arr[0]
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
