package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is ccc@gmail.com
email is abc@def.org
emai is aaaaqq@aaa.com.aa
`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
