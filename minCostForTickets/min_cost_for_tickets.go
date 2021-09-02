package minCostForTickets

import (
	"github.com/kevin-untrojb/interview-problems/utils"
	"math"
)

func mincostTickets(days []int, costs []int) int {
	return BacktrackminCostTickets(0, 0, 0, days, costs)
}

func BacktrackminCostTickets(actualIndex int, actualCost int, dayUntil int, days []int, costs []int) int {
	if actualIndex == len(days) {
		return actualCost
	}

	if dayUntil >= days[actualIndex] {
		return BacktrackminCostTickets(actualIndex+1, actualCost, dayUntil, days, costs)
	}
	costFor1daymore := BacktrackminCostTickets(actualIndex+1, actualCost+costs[0], dayUntil+1, days, costs)
	costFor7daymore := BacktrackminCostTickets(actualIndex+1, actualCost+costs[1], dayUntil+7, days, costs)
	costFor15DaysMore := BacktrackminCostTickets(actualIndex+1, actualCost+costs[2], dayUntil+15, days, costs)

	return min3(costFor1daymore, costFor7daymore, costFor15DaysMore)
}

func MincostTicketsDp(days []int, costs []int) int {
	opt := make([]int, len(days))

	for i := range opt {
		opt[i] = math.MaxInt32
	}
	prevDayCost := 0
	for i, day := range days {
		costFor1day := prevDayCost + costs[0]
		costFor7day := prevDayCost + costs[1]
		costFor15day := prevDayCost + costs[2]

		for j := i + 1; j < len(days); j++ {
			if days[j] > day+6 {
				break
			}
			opt[j] = utils.Min(opt[j], costFor7day)
		}
		for j := i + 1; j < len(days); j++ {
			if days[j] > day+29 {
				break
			}
			opt[j] = utils.Min(opt[j], costFor15day)
		}
		prevDayCost = min3(costFor1day, costFor7day, costFor15day)
		opt[i], prevDayCost = utils.Min(opt[i], prevDayCost), utils.Min(opt[i], prevDayCost)
	}
	return opt[len(days)-1]
}
