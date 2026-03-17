package main

import (
	"fmt"
)

// 1
type Person struct {
	Name string
	Age  int
}

// 2
type Book struct {
	Title, Author string
	Pages         int
}

// 3
type Car struct {
	Brand string
	Year  int
}

//4

// 5
type Student struct {
	Name  string
	Grade int
}

// 6
type Cirlce struct {
	Radius float64
}

func main() {
	//1
	user1 := Person{
		Name: "Tom",
		Age:  22}
	fmt.Println(user1.Name, user1.Age)
	//2
	book1 := Book{
		Title:  "DS",
		Author: "Hayato",
		Pages:  100,
	}
	fmt.Println(book1)
	//3
	car := Car{
		Brand: "BMW",
		Year:  26,
	}
	carP := &car
	carP.Year = 25
	fmt.Println(carP)

	//4

	//5
	student1 := Student{
		Name:  "Tim",
		Grade: 5,
	}
	student2 := Student{
		Name:  "Bob",
		Grade: 4,
	}
	if student1.Grade > student2.Grade {
		fmt.Printf("Студент %s получил %d\n", student1.Name, student1.Grade)
	} else {
		fmt.Printf("Студент %s получил %d\n", student2.Name, student2.Grade)
	}
	//6
	Cirlce1 := Cirlce{
		Radius: 50,
	}
	fmt.Println(3.14 * Cirlce1.Radius * Cirlce1.Radius)

	user := Person{Name: "Bab", Age: 25}
	fmt.Println(user)
}
func (u Person) incrementAge(addAge int) {
	u.Age += addAge
	fmt.Println(u.Name, u.Age)
}
