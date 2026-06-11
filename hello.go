// Главные пакет программы должен называться как "main"
package main

import (
	"fmt"
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

func main() {
	words := "##Fourth Message@@@"

	words = strings.TrimPrefix(words, "##")

	if strings.HasSuffix(words, "@@@") {
		words = ""
	}

	fmt.Println(words)
}
