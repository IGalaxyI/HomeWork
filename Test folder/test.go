package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
	ex5()
	ex6()
	ex7()
	ex8()
	ex9()
}
func ex1() {
	temperature := rand.Intn(21) - 20
	fmt.Println(temperature)
	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature <= 20 && temperature >= 0 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}
	fmt.Println("Ex1 End")
}
func ex2() {
	score := rand.Intn(101)
	fmt.Println(score)
	if score >= 90 {
		fmt.Println("Отлично")
	} else if 70 <= score && 89 >= score {
		fmt.Println("Хорошо")
	} else if 50 <= score && 69 >= score {
		fmt.Println("Удовлетворительно")
	} else if 50 > score {
		fmt.Println("Не сдал")
	}
	fmt.Println("Ex2 End")
}

func ex3() {
	hour := rand.Intn(24)
	fmt.Println(hour)
	switch {
	case hour >= 0 && hour <= 5:
		fmt.Println("Ночь")
	case hour >= 6 && hour <= 11:
		fmt.Println("Утро")
	case hour >= 12 && hour <= 17:
		fmt.Println("День")
	case hour >= 18 && hour <= 23:
		fmt.Println("Вечер")
	}
	fmt.Println("Ex3 End")
}

func ex4() {
	number := 0
	fmt.Scan(&number)
	if number&2 == 0 {
		fmt.Println("Чётное число")
	} else {
		fmt.Println("Не чётное")
	}
	fmt.Println("Ex4 End")
}

func ex5() {
	var day string
	fmt.Scan(&day)
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Будний день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной день")
	default:
		fmt.Println("Некорректный день")
	}
	fmt.Println("Ex5 End")
}

func ex6() {
	balance := rand.Intn(11) - 10
	if balance <= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}
	fmt.Println("Ex6 End")
}

func ex7() {
	age := rand.Intn(22)
	if 0 <= age && age >= 12 {
		fmt.Println("Ребёнок")
	} else if 13 <= age && age >= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}
	fmt.Println("Ex7 End")
}

func ex8() {
	var command string
	fmt.Scan(&command)
	switch command {
	case "start":
		fmt.Println("Start")
	case "stop":
		fmt.Println("Stop")
	case "restart":
		fmt.Println("Restart")
	default:
		fmt.Println("Неизвестная команда")
	}
	fmt.Println("Ex8 End")
}

func ex9() {
	garde := rand.Intn(5) + 1
	switch garde {
	case 5:
		fmt.Println("A")
	case 4:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 2:
		fmt.Println("F")
	}
	fmt.Println("Ex9 End")
}
