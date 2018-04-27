package main

import (
	"fmt"
	regx "regexp"
)

func main() {

	var (
		option   int
		filename string
	)

	const fileExtPattern = `\.(\w+)$`
	const filenamePattern = `[^/]+$`

	fileExtComp := regx.MustCompile(fileExtPattern)
	filenameComp := regx.MustCompile(filenamePattern)

	fmt.Println("1 for file extension, 2 for filename, 0 for quit")
	fmt.Scanf("%d", &option)
	if option == 0 {
		return
	}

	fmt.Println("Please enter filename string (quit for exit):")
	fmt.Scanf("%s", &filename)

	quit_flag, _ := regx.MatchString("quit", filename)
	for quit_flag != true && option != 0 {

		switch option {
		case 1:
			fmt.Println("File extension: " + fileExtComp.FindString(filename))
		case 2:
			fmt.Println("Filename: " + filenameComp.FindString(filename))
		default:
			fmt.Println("Unexpected option")
		}

		fmt.Println("1 for file extension, 2 for filename, 0 for quit")
		fmt.Scanf("%d", &option)
		if option == 0 {
			return
		}

		fmt.Println("Please enter filename string (quit for exit):")
		fmt.Scanf("%s", &filename)
		quit_flag, _ = regx.MatchString("quit", filename)
	}

	return
}
