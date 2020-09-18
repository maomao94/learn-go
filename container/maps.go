package container

import "fmt"

func main() {
	m := map[string]string{
		"name":    "hehanpeng",
		"course":  "golang",
		"site":    "hehanpeng",
		"quality": "notBad",
	}
	fmt.Println(m)

	m2 := make(map[string]int) // m2 == empty map
	var m3 map[string]string   // m3 == nil
	fmt.Println(m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if courseName, ok = m["cause"]; ok {
		fmt.Println(courseName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
	fmt.Println(m)
}

/**
map的key
1.map使用哈希表，必须可以比较相等
2.除了slice，map，function的内建类型都可以作为key
3.Struct类型不包含上述字段，也可以作为key
*/
