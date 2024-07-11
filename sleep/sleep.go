package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println(os.Args[1:])
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Tick")
		_, err := fmt.Fprintln(os.Stderr, "Tick")
		if err != nil {
			panic(err)
		}
	}
}
