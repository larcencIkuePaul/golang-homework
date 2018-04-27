package main

import (
	"fmt"
	"strings"
)

func main() {
	source := "s3://media-conv-dst/uploads/video/file/aaf4dc20-9846-417a-a942-a238fd80c5f4/"
	target := "s3://media-conv-dst/"
	fmt.Println(source)
	result := strings.TrimPrefix(source, target)
	fmt.Println(result)
}
