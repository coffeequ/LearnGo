// Главные пакет программы должен называться как "main"
package main

import (
	"fmt"
	"math"
	"strings"
) //Подключаем функциональность пакета fmt

func BaseOper() {
	//Объявление переменной
	var hello string = "Hello World"

	fmt.Println(hello)

	//Определение нескольких переменных
	var (
		name string = "Tom"
		age  int    = 123
	)

	fmt.Println(name)
	fmt.Println(age)

	//Краткое определение переменной. Тип определяется при присваивании значения
	//Можно использовать только внутри функции.Выводится автоматически из присваиваемого значения.
	nameNew := "123"
	fmt.Println("print new name: ", nameNew)

	//Iota-индентификаторы
	//iote - индентификаторы, увеличиваются с объявлением каждой новой константой.
	const (
		C = iota
		c1
		C2 = iota
	)

	//Функция Println имеет ряд ососбенностей.
	// При выводе в конец добавляется перевод строки.
	// При выводе нескольких значений они разделяются пробелом
	fmt.Println(C, name, age)

	//Выводит некоьорые значения или набор значений,
	// только не добавляет перевод строки или пробелма между разными значениями.
	// Грубо говорят вывод данных без форматирования
	fmt.Print()

	//В функцию передается строка форматирования. Строка форматирования содержит набор спецификаторов.
	//Каждый спецификатор представляет набор символов. которые интерпретируются определенным образом и предваряются знаком процента.
	fmt.Printf("Hello world \n")

	//При делении стоит обратить внимание на то, что если в операции учавствуют два целых числа, то результат деления будет округляться до целого числа, даже если результат присваивается переменной типа float32/64
	//Чтобы результат представлял вещественное число, один из операндов также должен представлять вещественное число:
	var value float32 = 10 / 4.0
	fmt.Println(value)
}

func Binary() {
	var a int = 2 << 2

	var b int = 16 >> 3

	fmt.Println("a: ", a)

	fmt.Println("b: ", b)
}

func Operation() {
	//= Присваивает левому операнду значение правого операнда
	var a int = 123

	//Присваивает левому операнду, который ранее не был определено, значение правого операнда
	b := 5
	fmt.Println("a: ", a, "b: ", b)
}

func arrayLearn() {
	//Стандартная инициализация массива
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	//Не явно указать размерность массива
	var test = [...]int{2, 10, 20, 30}
	fmt.Println(test)

	//Можем самостоятельно переопределять ключи и значения к ним
	var colors [3]string = [3]string{2: "blue", 0: "red", 1: "yellow"}
	fmt.Println(colors[1])

	//Многомерные массивы
	numbersNew := [3][2]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}

	fmt.Println(numbersNew)

	numbersNew[0] = [2]int{3, 4}

	fmt.Println(numbersNew)

	//Массив работает как значения, а не как ссылка
	var arr [3]int = [3]int{1, 2, 3}

	var arrDiff [3]int = arr

	arrDiff[0] = 123

	fmt.Println(arr, arrDiff)
}

func Palinom() {
	fmt.Print("Введите слово на проверку полинома: ")

	//scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	word = scanner.Text()
	// }

	//Массив байтов. Каждый символ представляет с байте.
	//https://habr.com/ru/articles/929076/
	var word string = "абоба"

	//Возвращает не длину строки, а количество занимаемых байтов
	//twoPtr := (len(word) / 2)

	var runeWord []rune = []rune(word)

	var lenWord int = len(runeWord)

	var ptrLast = lenWord - 1

	var isPoliom bool = true

	for i := 0; i < ptrLast; i++ {
		if runeWord[i] != runeWord[ptrLast] {
			isPoliom = false
			break
		}
		ptrLast--
	}
	if isPoliom {
		fmt.Print("Слово: ", word, " является полиномом")
	} else {
		fmt.Print("Слово: ", word, " не является полиномом")
	}
}

func loops() {
	//Классический цикл for (C - подобный синтаксис)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Аналог while

	isInfinity := true

	for isInfinity {
		//Логика выполнения
		isInfinity = false
	}

	//Бесконечный цикл
	for {
		fmt.Println("Work work work")
		break
	}

	//for range (Для обхода коллекций)
	arr := []string{"Go", "Python", "Java"}

	for _, val := range arr {
		fmt.Println(val)
	} // где i - индекс, val - значение элемента
}

func lessonsOne() {
	words := "##Fourth Message"

	words = strings.TrimPrefix(words, "##")

	if strings.HasSuffix(words, "@@@") {
		words = ""
	}

	fmt.Println(words)
}

func arrayLessons1312() {
	// countRates := 6

	// arr := [...]int{0, 0, 6, 0, 9, 8}

	// left := 0

	// right := 0

	// for left < countRates {
	// 	if arr[left] > 0 {
	// 		arr[right] = arr[left]
	// 		arr[left] = 0
	// 		right++
	// 	}
	// 	left++
	// }

	// fmt.Println(arr)

	countRates := 7

	arr := [7]int{1, 2, -1, 4, 5, -1, 6}

	elements := -1

	left := 0

	right := 0

	for left < countRates {
		if arr[left] != elements {
			arr[right] = arr[left]
			right++
		}
		left++
	}

	arr2 := arr[:right]

	fmt.Println(arr2)
}

func FindMaxVal() {
	lenArr := 5

	arr := [5]int{95, 87, 100, 92, 85}

	maxVal := arr[0]

	for i := 1; i < lenArr; i++ {
		if maxVal < arr[i] {
			maxVal = arr[i]
		}
	}

	fmt.Println(maxVal)
}

func findDuoValue() {
	lenArr := 4
	arr := [4]int{1, 3, 4, 11}

	val := 0

	indxOne := 0

	indxTwo := 1

	minVal := int(math.Abs(float64(arr[indxOne] - arr[indxTwo])))

	for i := 1; i < lenArr-1; i++ {

		val = int(math.Abs(float64(arr[i] - arr[i+1])))

		if val < minVal {
			indxOne = i
			indxTwo = i + 1
		}
	}
	fmt.Println(arr[indxOne], arr[indxTwo])
}

func lastEntryVal() {
	//Реализация через стек
	countContains := 9

	arr := [9]int{7, 3, 4, 1, 10, 11, 12, 19, 21}

	stack := []int{}

	for i := 0; i < countContains; i++ {
		if arr[i]%2 == 0 {
			stack = append(stack, arr[i])
		}
	}

	if len(stack) == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(stack[len(stack)-1])
	}

	//Реализация без стека

	lastVal := -1
	for i := 0; i < countContains; i++ {
		if arr[i]%2 == 0 {
			lastVal = arr[i]
		}
	}
	fmt.Println(lastVal)
}

func SearchFirstIncludeValue() {
	lenArr := 5
	arr := [5]int{5, 7, 9, 11, 13}
	target := 20

	left := 0
	//Для lower_bound правая граница должна быть за последним элементом
	right := lenArr

	for left < right {
		middle := (left + right) / 2

		if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	fmt.Println("Target:", left)
}

func maxInt(valOne int, valTwo int) (res int) {
	if valOne > valTwo {
		return valOne
	} else {
		return valTwo
	}
}

func PetyaDevDiplom() {
	arrX := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	height := 2

	weight := 3

	countN := 6

	left := maxInt(height, weight)

	right := left * countN

	for left+1 < right {
		middle := (left + right) / 2

		plane := (arrX[middle] / height) * (arrX[middle] / weight)

		if plane >= countN {
			right = middle
		} else {
			left = middle
		}
	}

	fmt.Println(right)
}

func ternarnAndExponenc() {
	//len := 11

	arr := []int{8, 11, 12, 16, 18, 21, 33, 36, 48, 54, 63}

	target := 16

	border := 1

	lastEl := len(arr) - 1

	//Для хранения старого значения border => он же будет являться и левым указателем при бинарном поиске
	prev := 0

	//Экспоненциальный поиск
	for border < lastEl && arr[border] < target {
		prev = border
		border *= 2

		if border > lastEl {
			border = lastEl
		}
	}

	//Бинарный поиск
	left := prev

	right := border

	for left <= right {
		middle := (left + right) / 2

		if arr[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	fmt.Println(left, right)
}

func SelectionSort(arr []int) []int {

	lenArr := len(arr)

	for i := 0; i < lenArr-1; i++ {
		for j := i + 1; j < lenArr; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func InsertSort(arr []int) []int {
	count := len(arr)
	for i := 1; i < count; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
	return arr
}

func SelectionSortRepit(arr []int) []int {

	lenArr := len(arr)

	for i := 0; i < lenArr-1; i++ {
		for j := i + 1; j < lenArr; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{5, 4, 1, 2, 10, 100, 21, 33, 2, 10, 100, 21, 33, 2, 10, 100, 21, 33, 2, 10, 100, 21, 33}
	resArr := SelectionSortRepit(arr)
	fmt.Println(resArr)
}
