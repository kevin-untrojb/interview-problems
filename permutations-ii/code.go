package permutations_ii

import "sort"
//https://leetcode.com/problems/permutations-ii
func permuteUnique(nums []int) [][]int {
	visited := make([]bool,len(nums))
	actual :=make([]int,0)
	var solution [][]int
	for i:= range visited{
		visited[i]= false
	}
	sort.Ints(nums)
	_permuteUnique(nums,actual,visited,&solution)
	return solution
}

func _permuteUnique(nums []int, actual []int, visited []bool, solution* [][]int ){
	if (len(actual) == len(nums)) {
		*solution = append(*solution,copyPerm(actual) )
	}

	for i:=0; i<len(nums); i++{
		if visited[i]{
			continue
		}

		if i > 0 && nums[i] == nums[i - 1] && !visited[i - 1]{
			continue
		}


		visited[i] = true

		actual = append(actual,nums[i])
		_permuteUnique(nums,actual,visited,solution)
		visited[i]=false
		actual = actual[:len(actual)-1]
	}

}


func copyPerm(nums []int)[]int{
	cp := make([]int,len(nums))
	for i,v:= range nums{
		cp[i]= v
	}
	return cp
}