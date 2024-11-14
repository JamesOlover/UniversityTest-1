package main

import (
  "fmt"
)

// Определение структуры Auto
type Auto struct {
  Mark string // Mark автомобиля string -  Тип данных строка
  VIN    string // VIN-номер автомобиля string -  Тип данных строка
  Year    int    // Год производства int -  Тип данных число
}

func main() {
  // Создание экземпляра структуры Auto
  MyAuto := Auto{
    Mark: "Mercedes Brabus G800",
    VIN:   "W555-666-ASD-777",
    Year:   2024,
  }

  // Обрезаем марку до 20 символов
  if len(MyAuto.Mark) > 20 {
    MyAuto.Mark = MyAuto.Mark[:20]
  }

  // Обрезаем VIN-номер до 16 символов
  if len(MyAuto.VIN) > 16 {
    MyAuto.VIN = MyAuto.VIN[:16]
  }

  // Вывод информации об автомобиле
  fmt.Println("Mark:", MyAuto.Mark)
  fmt.Println("VIN-номер:", MyAuto.VIN)
  fmt.Println("Год производства:", MyAuto.Year)
}
