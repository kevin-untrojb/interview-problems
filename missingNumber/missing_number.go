package missingNumber

import "sort"

func TrivialMissingNumber(arr []int) int{
	n := len(arr) // busco la cantidad de N
	var temp int
	var i,j int
	// recorro por todos los valores de N incluido N
	for i = 0; i < n+1; i++{
		// el arreglo tiene n-1 numeros, ya que le falta 1
		for j = 0; j < n ; j++{
			// si encuentra uno que es igual, sale de este loop
			if i == arr[j]{
				temp = arr[j]
				break
			}
		}
		// pregunta si termino el loop por no encontrarlo o por encontrarlo
		if i != temp{
			// si no encontró el numero en el arrglo, es el que faltaba
			return i
		}
	}
	// si no encontro el valor, quiere decir que el que faltaba es el ultimo numero
	return n+1
}

func SortedMissingNumber(arr []int) int{
	var i int
	n := len(arr)
	// sort del array
	sort.Ints(arr)
	// recorro por todos los valores buscando el que falta
	for i = 0; i < n; i++{
		// si no lo ecuentro, lo devuelvo
		if i != arr[i]{
			return i
		}
	}
	// si estan todos los valores, es porque falta el ultimo
	return n
}

func HashMissingNumber(arr []int) int {
	var i int
	n := len(arr)
	m := make(map[int]int)
	// completo el mapa con los elementos
	for i= 0; i < len(arr) ; i++{
		m[arr[i]]++
	}

	// recorro todos los numeros buscando la clave que falta
	for i= 0; i < n ; i++{
		_,ok := m[i]
		if !ok{
			return i
		}
	}
	return n
}
func OptMissingNumber(arr []int) int {
	var i int
	n := len(arr)
	sum := 0
	// sumo todos los elementos
	for i= 0; i < len(arr) ; i++{
		sum += arr[i]
	}
	// calculo la suma total
	totalSum := n*(n+1)/2
	// devuelvo la diferencia
	return totalSum-sum
}
