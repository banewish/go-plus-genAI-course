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

func Greet(name string) string {
	if name == "" {
		return "Welcome, stranger! Welcome to Go."
	} else {
		return "Welcome, " + name + "! Welcome to Go.\n"
	}
}

func main() {
	fmt.Println(Greet(strings.TrimSpace(os.Args[1])))

}
