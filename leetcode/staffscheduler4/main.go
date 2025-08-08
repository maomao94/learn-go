package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// 时间区间
type TimeSlot struct {
	Start int // 分钟，从0开始
	End   int
}

// 任务结构，增加优先级字段
type Task struct {
	ID       string
	Name     string
	TimeSlot TimeSlot
	Priority int
}

// 员工任务分配
type TaskAssignment struct {
	Task     Task
	TimeSlot TimeSlot
}

// 员工结构，MaxConcurrent=1，无并发
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

// 判断员工是否可分配任务（无并发）
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
		{"T1", "巡视A", TimeSlot{0, 60}, 10},
		{"T2", "巡视B", TimeSlot{30, 90}, 8},
		{"T3", "巡视C", TimeSlot{60, 120}, 6},
		{"T4", "巡视D", TimeSlot{120, 150}, 9},
		{"T5", "巡视E", TimeSlot{150, 210}, 3},
		{"T6", "巡视F", TimeSlot{180, 240}, 7},
		{"T7", "巡视G", TimeSlot{210, 270}, 2},
	}

	employees := []*Employee{
		{"张三", nil, 1, 0},
		{"李四", nil, 1, 0},
		{"王五", nil, 1, 0},
	}

	// 按价值密度（优先级/时长）降序排序任务
	sort.Slice(tasks, func(i, j int) bool {
		lengthI := tasks[i].TimeSlot.End - tasks[i].TimeSlot.Start
		lengthJ := tasks[j].TimeSlot.End - tasks[j].TimeSlot.Start
		valI := float64(tasks[i].Priority) / float64(lengthI)
		valJ := float64(tasks[j].Priority) / float64(lengthJ)
		return valI > valJ
	})

	for _, task := range tasks {
		assigned := false
		for _, emp := range employees {
			if canAssign(emp, task.TimeSlot) {
				emp.Assignments = append(emp.Assignments, TaskAssignment{task, task.TimeSlot})
				emp.Load++
				fmt.Printf("任务【%s】(优先级%d, 时长%d) 分配给员工【%s】\n",
					task.Name, task.Priority, task.TimeSlot.End-task.TimeSlot.Start, emp.Name)
				assigned = true
				break
			}
		}
		if !assigned {
			fmt.Printf("任务【%s】无法分配，资源不足或时间冲突\n", task.Name)
		}
	}

	fmt.Println("\n排班时间占用表（单位：分钟，步长30）:")
	printScheduleTable(employees, 0, 300, 30)
}
