package main
import "fmt"

func main() {
  // Подается 10 целых положительных чисел в массив. Затем еще 3 пары чисел -
  // индексов элементов массива, которые нужно поменять местами. Вывести результат
  var (
    a, b int
    array [10]uint8 =       // объявление массива длинной 10 символов
     [10]uint8{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 } // заполнение массива с объвления
  )
/*
  for i := range(array) {   // итерация с помощью range по индексам
      fmt.Scan(&array[i])
  }
*/
  fmt.Println("Enter 3 pairs of indices to replace: ")
  for i := 0; i < 3; i++ {
      fmt.Scan(&a, &b)
      // меняем значения по введённым индексам
      array[a], array[b] = array[b], array[a]
  }

  for _, value := range(array) { // итерация с помощью range по значениям
      fmt.Printf("%d ", value)   // индексы не нужны по-этому вместо них: _
  }
  fmt.Print("\n")

  // На ввод в массив подаются пять целых чисел
  // Найти и вывести максимальное число в массиве
	var ( max = -100000
        arrayMax = [5]int{} // по дэфолту залоняется нулями
      )
  fmt.Println("Enter 5 numbers: ")
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		arrayMax[i] = a
        if arrayMax[i] > max {
            max = arrayMax[i]
        }
	}
  fmt.Println("Max number: ", max)

  /* Срезы(слайсы)
      Срез состоит из: указателя на базовый масив, длины и емкости -
      количество элементов между началом среза и концом базового массива
    var a []int - создание пустого среза
    a[i:j] - создаёт новый срез с i по j-1 индексов
    a = append(a, 4, 5) - добавляет елементы 4 и 5
    a = append(a[0:2], a[3:]...) - удаляет 2-ой елемент
    n := copy(b, a) - копирует срез а в b(если совпадает длина)
  */
  // Вводится n — количество элементов среза. Далее n чисел—элементов
  // Вывести количество положительных элементов в последовательности
  var n int
  sum := 0
  fmt.Println("Enter slice length: ")
	fmt.Scan(&n)
  slice := make([]int, n, n) // создание среза длинной и емкостью n

  fmt.Println("Enter slice elements: ")
	for i := range slice {
		fmt.Scan(&slice[i])
	}
  for i := 0; i < len(slice); i ++ { // узнать емкость - cap()
    if slice[i] > 0 { sum++ }
  }
  fmt.Println("Quantity positive: ", sum)
}
