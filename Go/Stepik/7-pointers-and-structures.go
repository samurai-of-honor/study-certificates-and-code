package main
import custom "fmt" // импорт c синонимом

func main() {
  var a, b int
  custom.Print("Enter two numbers: ")
  custom.Scan(&a, &b)
  mult(&a, &b)

  custom.Print("Enter two numbers: ")
  custom.Scan(&a, &b)
  replace(&a, &b)

  c := 5
  d := &c
  e := &d
  **e++ // указатель на указатель
  custom.Println("Double pointer:", c)

  custom.Print("Enter the amount of ammo and power: ")
  custom.Scan(&a, &b)
  testStruct := Person{true, a, b} // создание екземпляра(также через new)
  custom.Println(testStruct.Shoot()) // вызовы методов
  custom.Println(testStruct.RideBike())
  custom.Println(testStruct.Shoot())
  custom.Println(testStruct.RideBike())
}

// Написать функцию, которая умножает значения
// на которые ссылаются два указателя и выводит результат
func mult(x1 *int, x2 *int) {
  custom.Println("Multiplication result:", *x1 * *x2)
}

// Поменяйте местами значения переменных на которые ссылаются
// указатели. После этого переменные нужно вывести
func replace(x1 *int, x2 *int) {
  *x1, *x2 = *x2, *x1
  custom.Println("Swap result:", *x1, *x2)
}

/*
Необходимо реализовать структуру с полями On, Ammo и Power.
У неё должны быть методы: Shoot и RideBike, которые возвращают значение bool.
Если значение On == false, то оба метода вернут false.
Делать Shoot можно только при наличии Ammo (Ammo уменьшается на единицу,
а метод возвращает true), иначе вернет false. Метод RideBike работает также.
*/
type Person struct { // обьявление структуры
  On bool
  Ammo, Power int
}

func (p *Person) Shoot() bool { // создание метода
  if p.On {                   // доступ к полям структуры через .
    if p.Ammo > 0 {
      p.Ammo--
      return true
    }
  }
  return false
}

func (p *Person) RideBike() bool {
  if p.On {
    if p.Power > 0 {
      p.Power--
      return true
    }
  }
  return false
}
