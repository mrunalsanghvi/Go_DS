package playaround

import (
	"fmt"
)

//Returns a pointer to the slice of permutations of a given string
// it does so by creating a recursion tree and swapping each character as we progress

func Permutations(a []byte, l int, r int) {

	if l == r {
		fmt.Printf("%c\n", a)
	} else {
		// Permutations made
		for i := l; i <= r; i++ {

			// Swapping done
			a[i], a[l] = a[l], a[i]

			// Recursion called
			Permutations(a, l+1, r)

			//backtrack
			a[i], a[l] = a[l], a[i]
		}
	}
}

func Fib(value uint64, hMap map[uint64]uint64) uint64 {

	if item, ok := hMap[value]; ok {

		return item
	}

	item := (Fib(value-1, hMap) + Fib(value-2, hMap))
	hMap[value] = item
	return item
}

//Driver code
//iArr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
//	playaround.Permutations(iArr, 0, len(iArr)-1)
