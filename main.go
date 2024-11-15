package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	// args := []string{"."}
	for _, dir := range args {
		fmt.Printf("%s\n", dir)
		items := Read(dir)
		fmt.Printf("No. items %d\n", len(items))
		for _, item := range items {
			fmt.Printf("%s\n", item.Path)
			fmt.Printf("%s\n", item.Info[0])
		}
	}
}
