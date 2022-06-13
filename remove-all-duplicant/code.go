package remove_all_duplicant


func removeDuplicates(s string, k int) string {
	if s == ""{
		return s
	}

	count := 0
	index := 0
	actualChar :=s[0]
	for i := 0; i< len(s); i++{
		if count == k{
			return removeDuplicates(deleteElements(s,index,index+k),k)
		}

		if actualChar == s[i]{
			count++
			continue
		}

		count = 1
		index = i
		actualChar = s[i]
	}
	if count == k{
		return removeDuplicates(deleteElements(s,index,index+k),k)
	}
	
	return s
}
func deleteElements(s string, from int, to int ) string{
	return s[:from]+s[to:]
}
func find132pattern(nums []int) bool {
	accumulativePositiveDif := 0
	var previousMin,previousMax int
	twoMins:= true
	if len(nums) == 0 {
		return false
	}

	currentMin := nums[0]
	currentMax := nums[0]

	for i := 1; i < len(nums); i++ {
		delta := nums[i] - nums[i-1]

		if delta >= 0 {
			accumulativePositiveDif += delta

			if twoMins && nums[i] > previousMin && nums[i] < previousMax {
				return true
			}
		} else {
			if abs(delta) < accumulativePositiveDif{
				return true
			} else {
				if accumulativePositiveDif != 0 {
					accumulativePositiveDif = 0
					previousMin = currentMin
					previousMax = currentMax
					twoMins = true
				}
			}

		}

		if currentMin > nums[i]{
			currentMin = nums[i]
		}
			currentMin = nums[i]

		if currentMax < nums[i]{
			currentMax = nums[i]
		}

	}
	return false
}
func abs(i int) int{
	if i>0{
		return i
	}
	return -i
}