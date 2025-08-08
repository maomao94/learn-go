package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/emirpasic/gods/queues/priorityqueue"
)

// 时间区间
type TimeSlot struct {
	Start int
	End   int
}

// 任务
type Task struct {
	ID       string
	Name     string
	TimeSlot TimeSlot
	Priority int // 优先级，数字越大优先级越高
}

// 员工任务分配
type TaskAssignment struct {
	Task     Task
	TimeSlot TimeSlot
}

// 员工结构
type Employee struct {
	Name          string
	Assignments   []TaskAssignment
	MaxConcurrent int
	Load          int // 当前任务数量
}

// 判断两个时间区间是否重叠
func isOverlap(a, b TimeSlot) bool {
	return a.Start < b.End && b.Start < a.End
}

// 统计员工某时间段的任务并发数
func concurrentCount(emp *Employee, slot TimeSlot) int {
	count := 0
	for _, assign := range emp.Assignments {
		if isOverlap(assign.TimeSlot, slot) {
			count++
		}
	}
	return count
}

// 判断员工是否可分配任务
func canAssign(emp *Employee, slot TimeSlot) bool {
	return concurrentCount(emp, slot) < emp.MaxConcurrent
}

// 计算时间区间字符串，并返回最大宽度
func buildTimeIntervals(startTime, endTime, step int) ([]string, int) {
	intervals := []string{}
	maxLen := 0
	for t := startTime; t < endTime; t += step {
		s := fmt.Sprintf("%d~%d", t, t+step)
		intervals = append(intervals, s)
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}
	return intervals, maxLen
}

// 居中填充字符串，宽度为width
func centerPad(s string, width int) string {
	if len(s) >= width {
		return s
	}
	padding := width - len(s)
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}

// 打印排班时间占用表
func printScheduleTable(employees []*Employee, startTime, endTime, step int) {
	intervals, width := buildTimeIntervals(startTime, endTime, step)
	slots := len(intervals)

	// 打印时间轴
	fmt.Printf("%-8s", "员工")
	for _, interval := range intervals {
		fmt.Print("|")
		fmt.Print(centerPad(interval, width))
	}
	fmt.Println("|")

	// 打印分割线
	fmt.Printf("%s\n", strings.Repeat("-", 8+slots*(width+1)+1))

	for _, emp := range employees {
		row := make([]rune, slots)
		for i := range row {
			row[i] = ' '
		}
		for _, assign := range emp.Assignments {
			startIdx := (assign.TimeSlot.Start - startTime) / step
			endIdx := int(math.Ceil(float64(assign.TimeSlot.End-startTime) / float64(step)))

			if startIdx < 0 {
				startIdx = 0
			}
			if endIdx > slots {
				endIdx = slots
			}
			for i := startIdx; i < endIdx; i++ {
				row[i] = '█'
			}
		}

		fmt.Printf("%-8s", emp.Name)
		for _, c := range row {
			fmt.Print("|")
			if c == '█' {
				fmt.Print(strings.Repeat("█", width))
			} else {
				fmt.Print(strings.Repeat(" ", width))
			}
		}
		fmt.Println("|")
	}
}

func main() {
	tasks := []Task{
		{"T1", "巡检A", TimeSlot{0, 60}, 3},
		{"T2", "巡检B", TimeSlot{30, 90}, 2},
		{"T3", "巡检C", TimeSlot{30, 90}, 1},
		{"T4", "巡检D", TimeSlot{30, 90}, 4},
		{"T5", "巡检E", TimeSlot{0, 30}, 5},
		{"T6", "巡检F", TimeSlot{0, 30}, 2},
		{"T7", "巡检G", TimeSlot{0, 30}, 1},
		{"T8", "巡检H", TimeSlot{90, 120}, 3},
		{"T9", "门禁巡查", TimeSlot{60, 120}, 3},
		{"T10", "烟感测试", TimeSlot{60, 90}, 1},
		{"T11", "红外测试", TimeSlot{90, 150}, 2},
		{"T12", "深夜任务", TimeSlot{900, 930}, 5}, // 大数字测试
	}

	employees := []*Employee{
		{"张三", nil, 2, 0},
		{"李四", nil, 1, 0},
		{"王五", nil, 1, 0},
	}

	// 按任务优先级和时长算优先因子，排序任务(优先级/时长，越大越优先)
	sort.Slice(tasks, func(i, j int) bool {
		ti := tasks[i]
		tj := tasks[j]
		di := ti.TimeSlot.End - ti.TimeSlot.Start
		dj := tj.TimeSlot.End - tj.TimeSlot.Start
		vi := float64(ti.Priority) / float64(di)
		vj := float64(tj.Priority) / float64(dj)
		return vi > vj
	})

	// 创建优先队列，任务负载最少优先
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		e1 := a.(*Employee)
		e2 := b.(*Employee)
		return e1.Load - e2.Load
	})

	// 入队所有员工
	for _, emp := range employees {
		pq.Enqueue(emp)
	}

	for _, task := range tasks {
		assigned := false
		size := pq.Size()

		temp := []*Employee{}

		for i := 0; i < size; i++ {
			empRaw, _ := pq.Dequeue()
			emp := empRaw.(*Employee)

			if !assigned && canAssign(emp, task.TimeSlot) {
				emp.Assignments = append(emp.Assignments, TaskAssignment{task, task.TimeSlot})
				emp.Load++
				fmt.Printf("任务【%s】优先级%d分配给员工【%s】\n", task.Name, task.Priority, emp.Name)
				assigned = true
			}

			temp = append(temp, emp)
		}

		for _, e := range temp {
			pq.Enqueue(e)
		}

		if !assigned {
			fmt.Printf("任务【%s】优先级%d无法分配，资源不足或冲突\n", task.Name, task.Priority)
		}
	}

	fmt.Println("\n排班时间占用表（单位：分钟，步长30）:")
	printScheduleTable(employees, 0, 960, 30)
}
