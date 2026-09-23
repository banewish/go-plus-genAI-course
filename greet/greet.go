// Package greet містить логіку привітання користувача.
//
// Завдання 2: реалізуйте функцію Greet самостійно, вручну, без ШІ.
// Автоматичні тести дивіться у файлі greet_test.go.
package greet

import (
"fmt"
"strings"
)

func Greet(name string) string {
name = strings.TrimSpace(name)
if name == "" {
return "Hello, stranger! Welcome to Go."
}

// Capitalize each word for better formatting
name = strings.Title(strings.ToLower(name))

return fmt.Sprintf("Hello, %s! Welcome to Go.", name)
}
