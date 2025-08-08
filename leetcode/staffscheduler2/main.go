package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type TimeSlot struct {
	Start int
	End   int
}

type Task struct {
	ID       string
	Name     string
	TimeSlot TimeSlot
}

type TaskAssignment struct {
	Task     Task
	TimeSlot TimeSlot
}

type Employee struct {
	Name          string
	Assignments   []TaskAssignment
	MaxConcurrent int
}

func isOverlap(a, b TimeSlot) bool {
	return a.Start < b.End && b.Start < a.End
}

func concurrentCount(emp *Employee, slot TimeSlot) int {
	count := 0
	for _, assign := range emp.Assignments {
		if isOverlap(assign.TimeSlot, slot) {
			count++
		}
	}
	return count
}

func canAssign(emp *Employee, slot TimeSlot) bool {
	return concurrentCount(emp, slot) < emp.MaxConcurrent
}

func sortEmployeesByLoad(emps []Employee) {
	sort.Slice(emps, func(i, j int) bool {
		return len(emps[i].Assignments) < len(emps[j].Assignments)
	})
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

func printScheduleTable(employees []Employee, startTime, endTime, step int) {
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
		{"T1", "巡检A", TimeSlot{0, 60}},
		{"T2", "巡检B", TimeSlot{30, 90}},
		{"T2", "巡检C", TimeSlot{30, 90}},
		{"T2", "巡检D", TimeSlot{30, 90}},
		{"T2", "巡检E", TimeSlot{0, 30}},
		{"T2", "巡检F", TimeSlot{0, 30}},
		{"T2", "巡检G", TimeSlot{0, 30}},
		{"T3", "门禁巡查", TimeSlot{60, 120}},
		{"T4", "烟感测试", TimeSlot{60, 90}},
		{"T5", "红外测试", TimeSlot{90, 150}},
		{"T6", "深夜任务", TimeSlot{900, 930}}, // 大数字测试
	}

	employees := []Employee{
		{"张三", nil, 2},
		{"李四", nil, 1},
		{"王五", nil, 1},
	}

	for _, task := range tasks {
		sortEmployeesByLoad(employees)

		assigned := false
		for i := range employees {
			if canAssign(&employees[i], task.TimeSlot) {
				employees[i].Assignments = append(employees[i].Assignments, TaskAssignment{task, task.TimeSlot})
				fmt.Printf("任务【%s】分配给员工【%s】\n", task.Name, employees[i].Name)
				assigned = true
				break
			}
		}
		if !assigned {
			fmt.Printf("任务【%s】无法分配，资源不足或冲突\n", task.Name)
		}
	}

	fmt.Println("\n排班时间占用表（单位：分钟，步长30）:")
	printScheduleTable(employees, 0, 960, 30)
}
