package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println(os.Args[0])
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	for index, value := range os.Args {
		fmt.Printf("\n[%v] %v", index, value)
	}
	delay := time.Since(start)
	fmt.Printf("\n %v to execute this program", delay)

}
