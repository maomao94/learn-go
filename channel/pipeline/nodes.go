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

func Merge(int1, int2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-int1
		v2, ok2 := <-int2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-int1
			} else {
				out <- v2
				v2, ok2 = <-int2
			}
		}
		close(out)
	}()
	return out
}
