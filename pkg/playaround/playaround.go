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

//driver code for below
// func main() {
// 	arr := []int{2, 14, 0, 9, 8, 7, 6, 9, 7, 2, 13, 9, 18}
// 	fmt.Println(arr)
// 	arr = Sortdups(arr)

// 	fmt.Println(arr)
// }
func Sortdups(arr []int) (arr2 []int) {

	//input:[2 14 0 9 8 7 6 9 7 2 13 9 18]
	//output:[14 0 8 6 13 18 2 2 9 9 9 7 7] //sort duplicate elements at the end and do not lose the sequence of occurrence

	dupMap := make(map[int]int)
	//gets duplicates into map

	for idx, val := range arr {
		if _, ok := dupMap[val]; ok {
			continue
		}
		for idx2 := idx; idx2 < len(arr)-idx; idx2++ {

			if count, ok := dupMap[arr[idx2]]; ok { //check if the item exist, if not add and increment count to 1
				count += 1
				dupMap[arr[idx2]] = count
			} else {
				dupMap[arr[idx2]] = 1
			}
		}
	}
	//fmt.Println(dupMap)

	for _, val := range arr { //add non duplicate value insequence that they appear by a map lookup
		if val2 := dupMap[val]; val2 < 2 {
			arr2 = append(arr2, val)
			delete(dupMap, val)
		}
	}
	//fmt.Println(arr2)
	count := 0
	for _, val := range arr { //iterate over original array

		if val2, ok := dupMap[val]; ok {
			count++
			if count > len(dupMap) {
				break
			}
			for idx := 0; idx < val2; idx++ { //append to result array
				arr2 = append(arr2, val)
			}
		}
	}
	return
}

//Driver code
//iArr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
//	playaround.Permutations(iArr, 0, len(iArr)-1)
