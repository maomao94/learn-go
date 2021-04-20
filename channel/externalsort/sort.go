package main

import (
	"bufio"
	"fmt"
	"learn-go/channel/pipeline"
	"os"
)

func main() {
	p := createPipeline(
		"small.in", 512, 4)
	writrToFile(p, "small.out")
	printFile("small.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}

func writrToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	defer write.Flush()
	pipeline.WriteSink(write, p)
}

func createPipeline(filename string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults,
			pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}
