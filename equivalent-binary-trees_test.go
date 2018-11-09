package main

import (
  "golang.org/x/tour/tree"
  "testing"
)

func Test(t *testing.T) {
  var tests = []struct {
    k int
    want [10]int
  }{
    {1, [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
    {2, [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
  }
  var got [10]int
  for _, c := range tests {
    ch := make(chan int)
    go Walk(tree.New(c.k), ch)
    i := 0
    for d := range ch {
      got[i] = d
      i++
    }
    if got != c.want {
      t.Errorf("Walk(%v) == %v, want %v", c.k, got, c.want)
    }
  }

  if !Same(tree.New(1), tree.New(1)) {
    t.Error("New(1) should equal New(1)")
  }
	if Same(tree.New(1), tree.New(2)) {
    t.Error("New(1) should not equal New(2)")
  }
}

//	ch := make(chan int)
//	go Walk(tree.New(1), ch)
//	for i := range ch {
//		fmt.Println(i)
//	}
