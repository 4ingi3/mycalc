package input

import (
	"fmt"
)

// Функция для чтения строки с консоли с выводом приглашения
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
