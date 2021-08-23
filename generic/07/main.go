package main

import "fmt"

type addable interface {
	comparable
}
type set[T addable] map[T]struct{}

func (s set[T]) add(v T) { s[v] = struct{}{} }
func (s set[T]) contains(v T) bool {
	_, ok := s[v]
	return ok
}
func (s set[T]) len() int   { return len(s) }
func (s set[T]) delete(v T) { delete(s, v) }
func (s set[T]) iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}
func print[T addable](s T) { fmt.Printf("%v ", s) }

// 泛型Set类型（泛型方法）
// go run -gcflags=-G=3 main.go
func main() {
	s := make(set[string])
	s.add("红烧肉")
	s.add("清蒸鱼")
	s.add("九转大肠")
	s.add("大闸蟹")
	s.add("烤羊排")
	fmt.Printf("%v\n", s)
	if s.contains("大闸蟹") {
		println("包含大闸蟹")
	} else {
		println("不包含大闸蟹")
	}
	fmt.Printf("the len of set: %d\n", s.len())

	s.delete("大闸蟹")
	fmt.Println("\nafter delete 大闸蟹:")
	if s.contains("大闸蟹") {
		println("包含大闸蟹")
	} else {
		println("不包含大闸蟹")
	}
	fmt.Printf("the len of set: %d\n", s.len())
	s.iterate(func(x string) { fmt.Println("您点的菜: " + x) })
}
