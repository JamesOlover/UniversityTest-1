package main

import (
  "fmt"
)

func main() {
  var n1 int
  var n2 int
  fmt.Scan(&n1)
  fmt.Scan(&n2)
  result := (n1 + n2) / 2
  fmt.Printf("(%d + %d) / 2 = %d", n1, n2, result)
  
}
