package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func GoWalk(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go GoWalk(t1, ch1)
	go GoWalk(t2, ch2)
	for {
		val1, ok1 := <- ch1
		val2, ok2 := <- ch2
		if ok1 != ok2 || val1 != val2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go GoWalk(tree.New(1), ch)
	res1 := Same(tree.New(1), tree.New(1))
	res2 := Same(tree.New(1), tree.New(2))
	fmt.Printf("%t %t", res1, res2)
}
