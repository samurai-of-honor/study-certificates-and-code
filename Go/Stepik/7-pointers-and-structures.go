package main
import custom "fmt" // импорт c синонимом

func main() {
  var a, b int
  custom.Scan(&a, &b)
  mult(&a, &b)

  custom.Scan(&a, &b)
  replace(&a, &b)

  a = 5
  b = &a
  c := &b
  **c++ // указатель на указатель
  custom.Println(a)
}

func mult(x1 *int, x2 *int) {
  custom.Println(*x1 * *x2)
}

func replace(x1 *int, x2 *int) {
    *x1, *x2 = *x2, *x1
    custom.Println(*x1, *x2)
}
