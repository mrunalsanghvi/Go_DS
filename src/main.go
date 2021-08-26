package main

import (
	"fmt"

	"github.com/mrunalsanghvi/Go_DS/pkg/hello"
	"github.com/mrunalsanghvi/Go_DS/pkg/tree/bst"
)

func main() {
	fmt.Println(hello.Helloworld())
	var root *bst.Node
	bst.Insert(&root, 20)
	bst.Insert(&root, 30)
	bst.Insert(&root, 50)
	bst.Insert(&root, 40)
	bst.Insert(&root, 10)
	bst.Insert(&root, 7)
	bst.Insert(&root, 70)
	bst.Insert(&root, 500)

	fmt.Println(bst.Find(root, 10))

	bst.Traverse_preorder(root)
	root = bst.Delete(root, 50)
	fmt.Println("")
	bst.Traverse_inorder(root)
	fmt.Println("")
	bst.Traverse_postorder(root)
	fmt.Println("")
}
