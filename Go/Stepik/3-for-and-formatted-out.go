// решения задач не оптимальны и лишь отражают функционал пройденного материала
package main
import "fmt"

func main() {
    var a, b int

    // C клавиатуры считываются два натуральных числа A и B
    // Вывести сумму всех чисел от A до B
    var sum int
    fmt.Println("Enter a, then b(a < b): ")
    fmt.Scan(&a, &b)
    for a <= b {
        sum += a
        a++
    }
    fmt.Println(sum)

    // Последовательность натуральных чисел завершается числом 0
    // Определить количество элементов, которые равны наибольшему элементу
    counter := 0
    max := 0
    fmt.Println("Enter nums(if you type 0 input will stop): ")
    for fmt.Scan(&a); a != 0; fmt.Scan(&a) { // ввод чисел пока не встретится 0
        if a > max {
            max = a
            counter = 1
        } else if a == max {
            counter++
        }
    }
    fmt.Println(counter)

    // Cчитывает целые числа с консоли
    // 1. если число меньше 10, то пропускаем
    // 2. если число больше 100, то прекращаем считывать числа
    // 3. в остальных случаях вывести это число в консоль
    fmt.Println("Enter nums(if you type 0 or nums > 100 input will stop): ")
    for fmt.Scan(&a); a != 0; fmt.Scan(&a){
        if a > 100 {
            break           // выход из цикла
        } else if a < 10 {
            continue        // следующая итерация цикла
        } else {
            fmt.Println(a)
        }
    }

    // Вклад в банке составляет X $
    // Ежегодно он увеличивается на P процентов
    // Определить, через сколько лет вклад составит не менее Y $
    var x, p, y float32
    var year int = 0
    fmt.Println("Enter deposit, then percent, then required amount: ")
    fmt.Scan(&x, &p, &y)
    for {                 // бесконечный цикл
        if x >= y {
            fmt.Println("The required amount is collected through", year, "year(s)")
            break
        } else {
            x *= ((100 + p) / 100)
            year++
        }
    }
}
