package main

import(
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func (t *tree.Tree) {
        if (t == nil) {
            return
        }
        walker(t.Left)
        ch <- t.Value
        walker(t.Right)
    }
    walker(t)
    close(ch)
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <- ch1
		v2, ok2 := <- ch2

//		fmt.Printf("v1 = %v, v2 = %v, ok1 = %v, ok2 = %v\n", v1, v2, ok1, ok2)
		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 {
			break
		}
	}
	return true
}

func main() {
//	ch := make(chan int)
//	go Walk(tree.New(1), ch)
//	for i := range ch {
//		fmt.Println(i)
//	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	return
}
