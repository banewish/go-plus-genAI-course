package greet

import (
	"fmt"
	"os"
	"strings"
)

func Greet(name string) {
	
	if name == "" {
		fmt.Println("Welcome, stranger! Welcome to Go.")
	} else {
		fmt.Printf("Welcome, %s! Welcome to Go.\n", name)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lesson1 <name>")
		return
	}

	Greet(strings.TrimSpace(os.Args[1]))
}
