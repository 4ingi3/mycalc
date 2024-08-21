package calc

// Получаем строку без кавычек
func GetStringArg(s string) string {
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1] // Возвращаем строку без первой и последней кавычки
	}
	return ""
}

// Преобразуем строку в целое число
func StringToInt(s string) (int, bool) {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' { // Проверяем, является ли символ цифрой
			return 0, false
		}
		result = result*10 + int(s[i]-'0') // Добавляем цифру к результату
	}
	return result, true
}

// Повторяем строку n раз
func RepeatString(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s // Добавляем строку к результату n раз
	}
	return result
}

// Делим строку на n частей
func DivideString(s string, n int) string {
	if n > 0 && len(s) >= n {
		return s[:len(s)/n] // Возвращаем первую часть строки
	}
	return ""
}

// Удаляем подстроку из строки
func RemoveSubstring(s, sub string) string {
	result := ""
	i := 0
	for i < len(s) {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub { // Если находим подстроку
			i += len(sub) // Пропускаем подстроку
		} else {
			result += string(s[i]) // Добавляем символ к результату
			i++
		}
	}
	return result
}
