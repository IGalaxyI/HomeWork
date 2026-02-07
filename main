package main

import "fmt"

func main() {
	//1
	run := [...]string{"Бег", "Берпи"}
	walk := [...]string{"Ходьба", "Ходьба в горку"}

	fmt.Println("Бег", len(run))
	fmt.Println("Ходьба", len(walk))

	//2
	subjectsList := [...]string{"Химия", "География", "Физика"}
	fmt.Println(subjectsList[0], subjectsList[len(subjectsList)-1])
	subjectsList[1] = "Биология"
	fmt.Println(subjectsList)

	//3
	M := [...]string{"Tom", "35", "New York"}
	name := M[0]
	age := M[1]
	home := M[2]
	fmt.Println(name, age, home)

	//4
	numbersList := [...]int{1, 2, 3, 4, 5}
	s := 0
	for i := 0; i < len(numbersList); i++ {
		if numbersList[i] == 3 {
			s++
		}
	}
	if s == 1 {
		fmt.Println("Число 3 есть в массиве")
	} else {
		fmt.Println("Число 3 отсутвует в массиве")
	}

	//5
	friendList := [...]string{"Александр", "Никита"}
	point := false
	for _, friend := range friendList {
		if friend == "Bekbolat" {
			point = true
		}
	}
	if point == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	//6
	firstList := [...]int{1, 2, 3}
	secondList := [...]int{1, 2, 3}
	count := 0
	for i := 0; i < len(firstList); i++ {
		if firstList == secondList {
			count++
		}
	}
	if count == len(firstList) {
		fmt.Println("Массивы равны")
	} else {
		fmt.Println("Массивы не равны")
	}

	//7
	myWishList := [...]string{"Скрепка", "Пакет БМВ"}
	friendsWishList := [...]string{"Булавка", "Транбалон"}
	registrationList := [len(myWishList) + len(friendsWishList)]string{}
	index := 0
	for j := 0; j < len(myWishList); j++ {
		registrationList[index] = myWishList[j]
		index++
	}
	for j := 0; j < len(friendsWishList); j++ {
		registrationList[index] = friendsWishList[j]
		index++
	}
	fmt.Println(registrationList)

	//8
	toyList := [...]string{"car", "doll", "ball"}
	testToyList := toyList
	testToyList[1] = "boat"
	fmt.Println(toyList)
	fmt.Println(testToyList)
}
