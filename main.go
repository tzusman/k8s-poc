package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for true {
		fmt.Printf("ENV: [%s]\n", os.Getenv("TEST_VAR"))
		fmt.Printf("Secret: [%s]\n", os.Getenv("TEST_SECRET"))
		time.Sleep(500 * time.Millisecond)
	}
}
