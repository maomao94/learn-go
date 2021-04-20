package pipeline

import "sort"

func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		// sort
		sort.Ints(a)

		//Output
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}
