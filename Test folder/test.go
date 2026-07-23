package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}
func ex1() {
	//t
	fmt.Println(5 == 5)
	//t
	fmt.Println(10 != 3)
	//f
	fmt.Println(7 > 12)
	//t
	fmt.Println(15 < 20)
	//t
	fmt.Println(8 >= 8)
	//f
	fmt.Println(6 <= 4)
	//f
	fmt.Println((10 > 5) && (3 < 1))
	//t
	fmt.Println((10 > 5) || (3 < 1))
	//f
	fmt.Println(!(5 == 5))
	//t
	fmt.Println(!(7 < 3))
	//f
	fmt.Println(true && false)
	//f
	fmt.Println(false || false)
	//t
	fmt.Println(true || false)
	//t
	fmt.Println((4+6 == 10) && (9 > 2))
	//t
	fmt.Println((12/3 == 4) || (8 < 5))
	fmt.Println("ex1 End")
}
func ex2() {
	age := 24
	hasTicket := true
	canEnter := age >= 18 && hasTicket == true
	fmt.Println(canEnter)
	fmt.Println("ex2 End")
}

func ex3() {
	isLoggedIn := true
	isAdmin := true
	hasAccess := (isLoggedIn && isAdmin) || (isLoggedIn && !isAdmin)
	fmt.Println(hasAccess)
	fmt.Println("ex3 End")
}
