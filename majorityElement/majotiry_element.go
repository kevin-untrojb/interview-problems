package majorityElementfunc

import "sort"

func TrivialMajorityElement(arr []int) int {
	maxCount := 0
	index := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count > maxCount {
			index = i
			maxCount = count
		}
	}
	if maxCount > n/2 {
		return arr[index]
	}
	return -1
}

func SortedMajorityElement(arr []int) int{
	count := 0
	index := 0
	n := len(arr)

	sort.Ints(arr)
	for i := 1; i < n; i++{
		if arr[i]== arr[index]{
			count++
		} else {
			if count > n/2{
				return arr[index]
			}
			count = 0
			index = i
		}
	}
	return -1
}

func HashMajorityElement(arr []int) int{
	n := len(arr)
	m := make(map[int]int)
	// completo el mapa con los y le sumo 1 al contador de variables
	for i:= 0; i < n ; i++{
		m[arr[i]]++
	}

	// busco si existe un contador mayor a n/2
	for i:= 0; i < n ; i++{
		count ,ok := m[i]
		if ok && count > n/2 {
			return arr[i]
		}
	}
	return -1
}

func OptimalMajorityElement(arr []int) int{
	count := 1
	index := 0
	n := len(arr)
	// busco el elemento que mas se repite en el arreglo
	for i := 1; i < n; i++{
		if arr[index] == arr[i]{
			count ++
		}else {
			count --
		}
		if count == 0{
			count = 1
			index = i
		}
	}

	// analizo si se repite mas de n/2 veces
	count = 0
	for i := 0; i < n; i++{
		if arr[index] == arr[i]{
			count++
		}
	}
	if count > n/2 {
		return arr[index]
	}
	return -1
}

