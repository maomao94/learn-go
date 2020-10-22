package loop

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/**
省略初始条件，相当于while
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/**
省略初始条件，递增条件
*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}

func printFileContents(read io.Reader) {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 ---> 1101
		convertToBin(0))
	printFile("abc.txt")
	s := `klkl
asasa
sasa
sasa
sasa

sasa`
	printFileContents(strings.NewReader(s))

	//forever()
}

/**
for的条件不需要括号
for的条件可以省略初始条件，结束条件，递增表达式
*/
