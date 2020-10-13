package main

import (
	"fmt"
	"os"

	"github.com/adityaputra11/mojek/cmd"
)

func init() {
	fmt.Println("Init")
}
func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
