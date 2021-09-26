package bst

import (
	"fmt"
)

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
			if getHeight(t.left) < getHeight(t.right) {
				tmp := findMin(t.left)
				t.data = tmp.data
				t.left = Delete(t.left, tmp.data)
			} else {
				tmp := findMax(t.left)
				t.data = tmp.data
				t.left = Delete(t.left, tmp.data)
			}
		}
	}
	return t
}

func getHeight(t *Node) int {
	if getHeightLeft(t) > getHeightRight(t) {
		return getHeightLeft(t)
	} else {
		return getHeightRight(t)
	}
}
func getHeightLeft(t *Node) int {
	leftheight := 0
	if t == nil {
		return 0
	}

	leftheight = getHeightLeft(t.left) + 1

	return leftheight

}

func getHeightRight(t *Node) int {
	rightheight := 0
	if t == nil {
		return 0
	}

	rightheight = getHeightLeft(t.right) + 1

	return rightheight
}

func findMax(t *Node) *Node {

	for t.right != nil {
		t = t.right
	}

	return t
}

func findMin(t *Node) *Node {

	for t.left != nil {
		t = t.left
	}

	return t
}

var result [][]uint64

func preorderTraversal(t *Node, iArr *[]uint64) {
	if t == nil {
		return
	}
	*iArr = append(*iArr, t.data)
	preorderTraversal(t.left, iArr)
	//*iArr = append(*iArr, t.data)
	preorderTraversal(t.right, iArr)
	//*iArr = append(*iArr, t.data)
	return
}

//Tree Driver code
// var root *bst.Node
// bst.Insert(&root, 20)
// bst.Insert(&root, 30)
// bst.Insert(&root, 50)
// bst.Insert(&root, 40)
// bst.Insert(&root, 10)
// bst.Insert(&root, 7)
// bst.Insert(&root, 70)
// bst.Insert(&root, 500)

// fmt.Println(bst.Find(root, 10))

// bst.Traverse_preorder(root)
// root = bst.Delete(root, 50)
// fmt.Println("")
// bst.Traverse_inorder(root)
// fmt.Println("")
// bst.Traverse_postorder(root)
