package main

import (
	"fmt"

	sr "github.com/tuvistavie/securerandom"
)

func main() {
	uuid, _ := sr.Uuid()
	fmt.Println("Uuid: " + uuid)
}
