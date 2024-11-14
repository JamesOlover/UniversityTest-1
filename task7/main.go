package main

import (
  "fmt"
)

func main() {
  var n, m int

  // Считываем размерность массива из консоли
  fmt.Print("Введите количество строк (n): ")
  fmt.Scan(&n)
  fmt.Print("Введите количество столбцов (m): ")
  fmt.Scan(&m)

  // Выделяем память под двумерный массив
  array := make([][]float64, n)
  for i := range array { // i - каждый элемент массива. пробегается по каждому элем массива
    array[i] = make([]float64, m)
  }

  // Заполняем массив числами от 1 до n*m
  counter := 1
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      array[i][j] = float64(counter)
      counter++
    }
  }

  // Выводим массив на экран
  fmt.Println("Массив:")
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      fmt.Printf("%.0f ", array[i][j]) // %.0f - флаг для выводв МАССИВА  i-строки m-столбцы 
    }
    fmt.Println() // чтоб столбцы были 
  }

  // В Go память автоматически освобождается сборщиком мусора,
  // поэтому явно освобождать память не нужно.
}
