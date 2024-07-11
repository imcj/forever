package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Tick")
		_, err := fmt.Fprintln(os.Stderr, "Tick")
		if err != nil {
			panic(err)
		}
	}
}
