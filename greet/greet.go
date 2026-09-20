// Package greet містить логіку привітання користувача.
//
// Завдання 2: реалізуйте функцію Greet самостійно, вручну, без ШІ.
// Автоматичні тести дивіться у файлі greet_test.go.
package greet

import (
	"strings"
)

func Greet(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello, stranger! Welcome to Go."
	} else {
		return "Hello, " + name + "! Welcome to Go."
	}
}
