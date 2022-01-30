package main
import . "fmt" // . - импортируем все функции пакета

func main() {
// Дано трехзначное число. Найти сумму его цифр.
    var a, b, c int
    Println("Enter 3-digit number: ")
    Scanf("%1d%1d%1d", &a, &b, &c)
    Println("Sum:", a + b + c)

    Println("Enter any number:")
    Scan(&a)
    Println("Powers of two to number:")
    for i := 1; i <= a; i *= 2 {
        Print(i, " ")
    }

// Идёт k-я секунда суток. Определить, сколько часов и минут прошло с начала суток.
    var second, h, m uint32
    Println("\nEnter any second of the day: ")
    Scan(&second)
    h = second / 3600
    m = (second - (h*3600)) / 60
    Printf("It is %d hours %d minutes.\n", h, m)

// Нужно проверить, является ли треугольник прямоугольным.
    Println("Enter katet, katet, hypotenuse: ")
    Scan(&a, &b, &c)

    if a * a + b * b == c * c {
      Println("Rectangular")
    } else {
      Println("Not rectangular")
    }

// Найти количество минимальных элементов в последовательности
    var count, n, min, sum int
    Println("Enter the length of the sequence and its elements: ")
    Scan(&count)

    for i := 0; i < count; i++ {
      Scan(&n)
      if min == 0 || n < min {
        min = n
        sum = 1
      } else if n == min {
         sum++
      }
    }
    Println("Minimum element count:", sum)

// Дано число n (0<n<100). Вывести n и одно из слов: korov, korova или korovy
    Print("Enter 0 < n < 100: ")
    Scan(&n)
    if n > 0 {
      switch {
        case n % 10 == 1 && n != 11 : Println(n, "korova")
        case n % 10 == 2 && n != 12 : Println(n, "korovy")
        case n % 10 == 3 && n != 13 : Println(n, "korovy")
        case n % 10 == 4 && n != 14 : Println(n, "korovy")
        default : Println(n, "korov")
      }
    }
}
