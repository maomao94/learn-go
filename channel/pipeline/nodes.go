package pipeline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

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

func ReaderSource(reader io.Reader, chuckSize int) <-chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil ||
				(chuckSize != -1 && bytesRead >= chuckSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func MergeN(intputs ...<-chan int) <-chan int {
	if len(intputs) == 1 {
		return intputs[0]
	}
	m := len(intputs) / 2
	// inputs[0..m) and inputs [m..end)
	return Merge(
		MergeN(intputs[:m]...),
		MergeN(intputs[m:]...))
}
