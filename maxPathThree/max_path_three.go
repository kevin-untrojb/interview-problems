package maxPathThree

import "github.com/kevin-untrojb/interview-problems/utils"

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func maxPathThree(T *Tree) int {
	m := make(map[int]int)
	return _solution(T, m)
}

func _solution(T *Tree, m map[int]int) int {
	if T == nil {
		return 0
	}

	_, ok := m[T.X]
	if ok {
		return 0
	}

	m[T.X]++
	MaxPathSons := utils.Max(_solution(T.L, m)+1, _solution(T.R, m)+1)

	m[T.X]--
	if m[T.X] == 0 {
		delete(m, T.X)
	}
	return MaxPathSons
}
