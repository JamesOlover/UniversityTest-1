package main

import (
  "fmt"
)
func min(a, b int)  { // функция принимает 2 аргумента типа число
  if a < b {
    fmt.Printf("Минимум: %d",a)
  } else {
    fmt.Printf("Минимум: %d",b)
  }
  
}

func main() {
  var num1, num2 int

  fmt.Print("Введите два целых числа: ")

  fmt.Scan(&num1, &num2) // Ввод чисел

  min(num1, num2)
}
