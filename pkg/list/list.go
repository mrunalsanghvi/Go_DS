package list

import (
	"fmt"
)

type Node struct {
	Next *Node
	data interface{}
}

func Reverse(N *Node) *Node {
	var rList *Node
	if N.Next == nil {
		return N
	}
	rList = Reverse(N.Next)

	N.Next.Next = N
	N.Next = nil
	return rList
}
func Add(N *Node, data interface{}) *Node {

	tmp := N
	if tmp == nil {
		N = new(Node)
		N.data = data
		return N
	}

	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = new(Node)
	tmp.Next.data = data
	return N

}

func SortedMerge(list1 *Node, list2 *Node) *Node {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.data.(int) < list2.data.(int) {
		list1.Next = SortedMerge(list1.Next, list2)
		return list1
	} else {
		list2.Next = SortedMerge(list1, list2.Next)
		return list2
	}
}

func PrintList(N *Node) {
	for N != nil {
		fmt.Printf("%v, ", N.data)
		N = N.Next
	}
	fmt.Printf("\n")
}

func IfIntersection(ls1 *Node, ls2 *Node) bool {
	tmp1 := ls1
	for tmp1.Next != nil {
		tmp1 = tmp1.Next
	}
	tmp2 := ls2
	for tmp2.Next != nil {
		tmp2 = tmp2.Next
	}
	if tmp2 == tmp1 {
		return true
	} else {
		return false
	}

}

// var ls *list.Node
// var ls2 *list.Node
// var ls3 *list.Node

// ls = list.Add(ls, 10)
// ls = list.Add(ls, 9)
// ls = list.Add(ls, 8)
// ls = list.Add(ls, 7)
// ls = list.Add(ls, 6)

// ls2 = list.Add(ls2, 1)
// ls2 = list.Add(ls2, 2)
// ls2 = list.Add(ls2, 3)
// ls2 = list.Add(ls2, 4)
// ls2 = list.Add(ls2, 5)

// ls = list.Reverse(ls)

// list.PrintList(ls)
// ls3 = list.SortedMerge(ls, ls2)
//	list.PrintList(ls3)
