package main

import "fmt"

func merge(arr *[]int, l int, m int, r int) {

	//arr1 := (*arr)[l : m+1]
	//arr2 := (*arr)[m+1 : r+1]
	var arr1 []int
	var arr2 []int

	l1, l2 := m-l+1, r-m
	for _, aval := range (*arr)[l : m+1] {
		arr1 = append(arr1, aval)
	}
	for _, aval := range (*arr)[m+1 : r+1] {
		arr2 = append(arr2, aval)
	}

	i, j, k := 0, 0, l
	//for i < (m-l+1) && j <= (r-m) {
	for (i < l1) && (j < l2) {
		if arr1[i] <= arr2[j] {
			(*arr)[k] = arr1[i]
			i++
		} else {

			(*arr)[k] = arr2[j]
			j++
		}
		k++
	}
	for i < l1 {
		(*arr)[k] = arr1[i]
		k++
		i++
	}
	for j < l2 {
		(*arr)[k] = arr2[j]
		k++
		j++
	}

	//fmt.Println(arr)
}

func merge_sort(arr *[]int, l int, r int) {

	if l >= r {
		return
	}
	med := l + (r-l)/2
	//fmt.Println(l, med, r)
	/*if l == 3 {
		return
	}*/
	merge_sort(arr, l, med)
	merge_sort(arr, med+1, r)
	merge(arr, l, med, r)

}

func main() {
	/*	e := sort.{
			l:  1,
			r:  10,
			sl: []int{10, 20, 30, 1, 4, 5, 6},
		}
		e.Merge_sort()*/
	/*arr := []int{2, 3, 4}
	func(arr *[]int) {
		//arr1 := (*arr)
		(*arr) = append(*arr, 1)

		fmt.Printf("%v\n", *arr)
	}(&arr)
	fmt.Printf("%v\n", arr)
	return*/
	arr := []int{4, 6, 5, 1, 2, 9, 8, 7, 3, 10}
	fmt.Println(arr)
	merge_sort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
	return
}
