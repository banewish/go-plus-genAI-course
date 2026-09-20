// Package greet містить логіку привітання користувача.
//
// Завдання 2: реалізуйте функцію Greet самостійно, вручну, без ШІ.
// Автоматичні тести дивіться у файлі greet_test.go.
package greet

import (
	"fmt"
	"os"
	"strings"
)

// This is what i wrote by hand
func Greet(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello, stranger! Welcome to Go."
	}
	return "Hello, " + name + "! Welcome to Go."
}

func main() {
	fmt.Println(Greet(strings.TrimSpace(os.Args[1])))
}
