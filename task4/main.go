package main

import (
  "fmt"
)

// Функция для проверки, является ли число трехзначным
func isThreeDigitNumber(number int) bool {
  // Проверяем, что число находится в диапазоне от 100 до 999
  return number >= 100 && number <= 999
}

func main() {
  var inputNumber int

  // Запрашиваем у пользователя ввод числа
  fmt.Print("Введите число: ")
  fmt.Scan(&inputNumber)

  // Проверяем, является ли введенное число трехзначным(8-10)
  result := isThreeDigitNumber(inputNumber)

  // Выводим результат
  fmt.Printf("Трехзначное число: %v", result)
}
