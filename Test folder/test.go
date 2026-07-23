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
}

func ex1() {
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}
	fmt.Println("Ex1 End")
}
func ex2() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println("Ex2 End")
}

func ex3() {
	number := rand.Intn(9) + 1
	fmt.Println(number)
	for i := 1; i < 10; i++ {
		fmt.Println(number * i)
	}
	fmt.Println("Ex3 End")
}

func ex4() {
	n := rand.Intn(100)
	s := []int{}
	fmt.Println(n)
	for i := 1; i <= n; i++ {
		v := i % 3
		if v == 0 {
			s = append(s, i)
		}
	}
	fmt.Println(s)
	fmt.Println("Ex4 End")
}

func ex5() {
	n := rand.Intn(1000)
	fmt.Println(n)
	sum := 0
	for i := 1; i < n; i++ {
		n /= 10
		sum += 1
	}
	fmt.Println(sum)
	fmt.Println("Ex5 End")
}

func ex6() {
	text := "TEST"
	for i := 0; i < len(text); i++ {
		fmt.Println(string(text[i]))
	}
	fmt.Println("Ex6 End")
}

func ex7() {
	balance := 3000
	for true {
		fmt.Println("Введите 1.2.3")
		a := 0
		fmt.Scan(&a)

		if a == 1 {
			fmt.Println(balance)
		} else if a == 2 {
			balance += 500
		} else if a == 3 {
			balance -= 200
		} else if a == 0 {
			break
		}
	}
	fmt.Println("Ex7 End")
}
