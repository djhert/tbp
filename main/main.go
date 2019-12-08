package main

import (
	"fmt"
	"os"
)

func main() {
	err := setup()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
