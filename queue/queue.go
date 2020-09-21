package queue

//别名方式扩展
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type QueueI []interface{}

func (q *QueueI) Push(v string) {
	*q = append(*q, v)
}

func (q *QueueI) Push1(v interface{}) {
	*q = append(*q, v.(string)) //限定类型
}

func (q *QueueI) Pop() string {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(string) //限定类型
}

func (q *QueueI) IsEmpty() bool {
	return len(*q) == 0
}
