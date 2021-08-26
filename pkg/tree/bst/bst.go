package bst

import (
	"fmt"
)

type comparator interface {
	compare(data *interface{}) int
}

type myData struct {
	h int
}

func (value *myData) compare(data *interface{}) int {
	datai := (*data).(myData)
	return value.h - datai.h
}

//func findInorderSuccessor(t *Node) *Node {
//
//	if t == nil {
//		return nil
//	}
//	if t.right != nil {
//		return findMin(t.right)
//	}4
//
//	return nil
//}

func IsCompare() {
	data := myData{
		h: 10,
	}
	var value comparator = &data
	var datai *interface{} = new(interface{})
	*datai = myData{h: 20}

	fmt.Println(value.compare(datai))
}

type Node struct {
	left  *Node
	right *Node
	data  uint64
}

func Insert(t **Node, value uint64) bool {
	res := false
	if *t == nil {
		*t = new(Node)
		(*t).data = value
		res = true
	} else {
		if value < (*t).data {
			res = Insert(&((*t).left), value)
		} else {
			res = Insert(&((*t).right), value)
		}
	}
	return res
}
func Find(root *Node, value uint64) bool {
	var res bool
	if root == nil {
		return false
	}
	if root.data > value && root.left != nil {
		res = Find(root.left, value)
	} else if root.data < value && root.right != nil {
		res = Find(root.right, value)
	} else {
		return true
	}
	return res
}

func Traverse_inorder(t *Node) {

	if t == nil {
		return
	}
	Traverse_inorder(t.left)
	fmt.Printf("%v, ", t.data)
	Traverse_inorder(t.right)
}

func Traverse_preorder(t *Node) {

	if t == nil {
		return
	}
	fmt.Printf("%v, ", t.data)
	Traverse_inorder(t.left)
	Traverse_inorder(t.right)
}
func Traverse_postorder(t *Node) {

	if t == nil {
		return
	}

	Traverse_inorder(t.left)
	Traverse_inorder(t.right)
	fmt.Printf("%v, ", t.data)
}

func Delete(t *Node, value uint64) *Node {

	if value < t.data {
		t.left = Delete(t.left, value)
	} else if value > t.data {
		t.right = Delete(t.right, value)
	} else {
		if t.right == nil {
			leftChild := t.left
			t = nil
			return leftChild
		} else if t.left == nil {
			rightChild := t.right
			t = nil
			return rightChild
		} else {
			tmp := findMin(t.left)
			t.data = tmp.data
			t.left = Delete(t.left, tmp.data)
		}
	}
	return t
}

func findMin(t *Node) *Node {

	for t.left != nil {
		t = t.left
	}

	return t
}
