package main
import . "fmt" // . - импортируем все функции пакета

func main() {
  var some_text string
  Println("Enter string: ")
  Scan(&some_text)
  fn(some_text)

  min := minimumFromFour() // присвоение возвращаемого результата в переменную
  Println("Minimum number:", min)

  res := vote(1, 0, 0)
  Println("More votes:", res)

  var n int
  Println("Enter Fibonacci element index:")
  Scan(&n)
  res = fib(n)
  Println(n, "element Fibonacci:", res)

  count, sum := sumInt(1, 2, 3, 4, 5, 6)
  Printf("Amount of numbers: %d\nSum: %d\n", count, sum)
}

// Написать функцию, которая будет принимать строку и выводить её
func fn(text string){
  Println("Your string:", text)
}

// Написать функцию, находящую наименьшее из четырех введённых чисел
func minimumFromFour() int { // тип возвращаемого значения(может быть несколько)
  var a, min int
  Println("Enter 4 numbers:")
  Scan(&min)
  for i := 1; i < 4; i++ {
      Scan(&a)
      if a < min {
          min = a
      }
  }
  return min
}

// Написать "функцию голосования", возвращающую 0 или 1,
// в зависимости от количества аргументов которые встречается чаще
func vote(x int, y int, z int) ( int ) {
  if x == y {
      return x
  }
  return z
}

// Вводится n. Необходимо вывести значение
// n-ого елемента последовательности Фибоначчи
func fib(n int) int {
  var a, b int = 1, 1
  for i := 1; i < n; i++ {
    a, b = b, a + b
  }
  return a
}

// Написать функцию принимающую переменное количество int аргументов,
// и возвращающую количество полученных аргументов и их сумму.
func sumInt(n ...int) (counter, sum int) {
	for _, elem := range n {
		counter++
    sum += elem
	}
  return counter, sum
}
