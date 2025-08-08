package main

import (
	"fmt"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

type Employee struct {
	Name     string
	TaskLoad int
}

type Task struct {
	ID   string
	Name string
}

func employeeComparator(a, b interface{}) int {
	e1 := a.(Employee)
	e2 := b.(Employee)
	return e1.TaskLoad - e2.TaskLoad
}

func main() {
	tasks := []Task{
		{"T1", "巡检 A"},
		{"T2", "巡检 B"},
		{"T3", "门禁巡查"},
		{"T4", "烟感测试"},
		{"T5", "红外测试"},
		{"T6", "摄像头检查"},
		{"T7", "电源检查"},
	}

	employees := []Employee{
		{"张三", 0},
		{"李四", 0},
		{"王五", 0},
	}

	pq := priorityqueue.NewWith(employeeComparator)

	// 员工入队
	for _, emp := range employees {
		pq.Enqueue(emp)
	}

	// 任务分配
	for _, task := range tasks {
		if pq.Empty() {
			fmt.Println("没有员工可以分配任务。")
			break
		}
		empRaw, _ := pq.Dequeue()
		emp := empRaw.(Employee)

		fmt.Printf("✅ 任务【%s】分配给员工【%s】（当前任务数: %d）\n", task.Name, emp.Name, emp.TaskLoad)

		emp.TaskLoad++
		pq.Enqueue(emp)
	}

	// 总结统计：从优先队列取出全部员工，打印最终任务数
	fmt.Println("\n📊 最终排班统计：")
	for !pq.Empty() {
		empRaw, _ := pq.Dequeue()
		emp := empRaw.(Employee)
		fmt.Printf("员工【%s】总任务数：%d\n", emp.Name, emp.TaskLoad)
	}
}
