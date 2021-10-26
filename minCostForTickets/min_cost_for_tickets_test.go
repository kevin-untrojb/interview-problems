package minCostForTickets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostTicketsBacktrack(t *testing.T) {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}

	assert.Equal(t, mincostTickets(days, costs), 11)
}

func TestMinCostTicketsDP(t *testing.T) {
	days := []int{1, 4, 6, 7, 8, 20}
	costs := []int{2, 7, 15}

	assert.Equal(t, MincostTicketsDp(days, costs), 11)
}
func TestMinCostTicketsDP2(t *testing.T) {
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs := []int{2, 7, 15}

	assert.Equal(t, MincostTicketsDp(days, costs), 17)
}

func TestMinCostTicketsDP3(t *testing.T) {
	days := []int{1, 4, 6, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22, 23, 27, 28}
	costs := []int{3, 13, 45}

	assert.Equal(t, MincostTicketsDp(days, costs), 44)
}

func TestMinCostTicketsWithDaysDP2(t *testing.T) {
	days := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}
	costs := []int{2, 7, 15}
	daysxCost := []int{1, 7, 15}

	assert.Equal(t, MincostTicketsWithDaysDp(days, costs, daysxCost), 17)
}
