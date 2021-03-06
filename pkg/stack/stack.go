package stack

import (
	"errors"
	"fmt"
)

type Stk struct {
	ptr  int64
	data [MAXSIZE]interface{}
}

const (
	MAXSIZE = 5
)

func Get() *Stk {
	return &(Stk{ptr: -1})

}

func Push(stackptr *Stk, data interface{}) (int, error) {
	if stackptr.ptr >= MAXSIZE-1 {
		return -1001, errors.New("stack is full")
	}
	stackptr.ptr++
	stackptr.data[stackptr.ptr] = data

	return 0, nil
}

func Pop(stackptr *Stk) (int, error) {
	if stackptr.ptr == -1 {
		return -1002, errors.New("stack is empty")
	}
	stackptr.data[stackptr.ptr] = nil
	stackptr.ptr--
	return 0, nil
}

func PrintStk(stackptr *Stk) {
	if stackptr.ptr > int64(-1) {
		for i := int64(0); i <= stackptr.ptr; i++ {
			fmt.Printf("%v", stackptr.data[i])
		}
	} else {
		fmt.Println("Stack is all Empty")
	}
}

//Stack Driver Code

// stackptr := stack.Get()
// i := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// for j := 0; j < 10; j++ {
// 	if _, ok := stack.Push(stackptr, i[j]); ok != nil {
// 		fmt.Println(ok)
// 		break
// 	}
// }
