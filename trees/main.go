package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker is a wrapper function for Walk to close the channel once done.
func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch) // This ensures that the channel is closed once Walk is done.
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walker(t1, ch1)
	go Walker(t2, ch2)

	for val := range ch1 {
		if val != <-ch2 {
			return false
		}
	}

	return true

}

func main() {
	t := tree.New(1)
	c := make(chan int)

	// Start Walk(t, ch) as a goroutine.
	go func() {
		Walk(t, c)
		close(c) // Closing the channel as we are done walking the tree.
	}()

	// Read values from the channel until it's closed.
	for v := range c {
		fmt.Println(v) // print the values received from the channel
	}
	t1, t2 := tree.New(1), tree.New(1)

	// Test with two trees which should be the same.
	same := Same(t1, t2)
	fmt.Printf("Should be true: %v\n", same)

	// Test with two trees which should be different.
	t3 := tree.New(2)
	notSame := Same(t1, t3)
	fmt.Printf("Should be false: %v\n", notSame)

}
