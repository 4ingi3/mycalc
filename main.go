package main

import (
	"fmt"
	"mycalc/calc"
	"mycalc/input"
)

func main() {
	// Считываем первую строку
	str1Input := input.ReadInput("Введите первую строку (в кавычках): ")
	str1 := calc.GetStringArg(str1Input)
	if str1 == "" {
		fmt.Println("Ошибка: первый аргумент должен быть строкой в кавычках.")
		return
	}

	// Считываем оператор
	operator := input.ReadInput("Введите оператор (+, -, *, /): ")

	var str2Input string
	var result string

	// Обработка ввода в зависимости от оператора
	if operator == "+" || operator == "-" {
		// Для операторов + и - мы ожидаем строку
		str2Input = input.ReadInput("Введите вторую строку (в кавычках): ")
		if str2 := calc.GetStringArg(str2Input); str2 != "" {
			if operator == "+" {
				result = str1 + str2 // Складываем две строки
			} else if operator == "-" {
				result = calc.RemoveSubstring(str1, str2) // Удаляем вторую строку из первой
			}
		} else {
			fmt.Println("Ошибка: третий аргумент должен быть строкой в кавычках.")
			return
		}
	} else if operator == "*" || operator == "/" {
		// Для операторов * и / мы ожидаем число
		str2Input = input.ReadInput("Введите число от 1 до 10: ")
		num, valid := calc.StringToInt(str2Input)
		if !valid || num < 1 || num > 10 {
			fmt.Println("Ошибка: третий аргумент должен быть числом от 1 до 10.")
			return
		}
		if operator == "*" {
			result = calc.RepeatString(str1, num) // Повторяем строку num раз
		} else if operator == "/" {
			result = calc.DivideString(str1, num) // Делим строку на num частей
		}
	} else {
		fmt.Println("Ошибка: неподдерживаемый оператор.")
		return
	}

	// Ограничиваем длину результата до 40 символов
	if len(result) > 40 {
		result = result[:40] + "..."
	}

	// Выводим результат
	fmt.Println(result)
}
