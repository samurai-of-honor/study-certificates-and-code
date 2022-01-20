package main

import ( // блочное обьявление
	"fmt"
)

const ( // блочное обьявление констант
	Sunday  = iota // 0 т.к. iota - значение индекса елемента
	Monday         // каждая константа копирует значение предыдущей
	Tuesday        // (по-этому все константы будут иметь значение iota)
	Wednesday
	Thursday // 4
	Friday   // 5
	Saturday // 6
)

func main() {
	fmt.Println(Sunday)   // вывод 0
	fmt.Println(Saturday) // вывод 6

  var a, year int
  // По данному трехзначному числу определить, все ли его цифры различны
  fmt.Print("Enter three-digit number: ")
	fmt.Scan(&a)
  a100 := a / 100
  a10 := (a % 100) / 10
  a1 := a % 10

  fmt.Print("All numbers are unique: ")
  switch {                       // условие не в switch a в cases
  case a1 == a10: fallthrough    // следующий case выполнится вне зависимости
  case a10 == a100: fallthrough  // от того истинно ли его (следующего case) условие
  case a1 == a100: fmt.Println("NO")
  default: fmt.Println("YES")
  }

  // Дано неотрицательное целое число. Найти и вывести первую цифру числа
  fmt.Print("Enter non-negative integer: ")
  fmt.Scan(&a)
  fmt.Print("First number: ")
	switch {
	case a < 10: fmt.Println(a)
	case a < 100: fmt.Println(a / 10)
	case a < 1000: fmt.Println(a / 100)
	case a < 10000: fmt.Println(a / 1000)
	case a < 100000: fmt.Println(a / 10000)
	}

  // Проверить в шестизначном номере совпадает ли сумма первых трёх цифр
  // с суммой трёх последних
  fmt.Print("Enter six-digit number: ")
  fmt.Scan(&a)
  first := a % 1000000 / 100000
  second := a % 100000 / 10000
  third := a % 10000 / 1000
  fourth := a % 1000 / 100
  fifth := a % 100 / 10
  sixth := a % 10

  fmt.Print("Sum of first three = sum of last: ")
  if (first + second + third) == (fourth + fifth + sixth) {
      fmt.Println("YES")
  } else {
      fmt.Println("NO")
  }

  // Определить, является ли введённый год високосным
  fmt.Print("Enter year: ")
  fmt.Scan(&year)
  fmt.Print("Leap year: ")
    if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
