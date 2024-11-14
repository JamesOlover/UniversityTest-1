package main

import (
  "fmt"
)

// Функция для обмена значениями двух булевых переменных
func swapBoolean(a *bool, b *bool) { // a and b хранят ССЫЛКУ на булевое значение а не просто true false
  // Используем временную переменную для обмена значениями
  temp := *a
  *a = *b
  *b = temp
}

func main() {
  x := true
  y := false

  fmt.Printf("До: x = %v, y = %v\n", x, y)

  // Передаем указатели на переменные x и y
  swapBoolean(&x, &y)

  fmt.Printf("После: x = %v, y = %v\n", x, y)
}
