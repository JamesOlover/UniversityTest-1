package main

import (
  "fmt"
)

// Функция для вычисления факториала
func factorial(n int) int {
  if n == 0 || n == 1 {  // || - или
    return 1
  }
  return n * factorial(n-1)
}

func main() {
  size := 10
  // Выделение памяти под динамический массив
  array := make([]int, size) // make - создаёт что то пустое. []int - массив с числами, size - максимальная длинна массива

  // Заполнение массива значениями факториалов индексов
  for i := 0; i < size; i++ {
    array[i] = factorial(i)
  }

  // Вывод массива на экран
  for i := 0; i < size; i++ {
    fmt.Printf("Элемент массива[%d] = %d\n", i, array[i])
  }
}
