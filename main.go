package main

import (
	"fmt"
)

func main() {
	// var users map[string]string
	// users = map[string]string{"Sanya": "12341254", "Nikita": "22", "Dimon": "23"}
	// fmt.Println(users)

	// result, ok := users["Sanya"]
	// fmt.Println(result, ok)

	// for _, v := range users {
	// 	if v == "1234" {
	// 		fmt.Printf("Too easy password User: \n")
	// 	} else {
	// 		fmt.Printf("User not found User: \n")
	// 	}
	// }

	// if ok {
	// 	if result == "1234" {
	// 		fmt.Println("Too easy password")
	// 	} else {
	// 		fmt.Println("User not found")
	// 	}
	// } else {
	// 	fmt.Println("Такого ключа нет в словаре")
	// }

	// fmt.Println(users["Sanya"])

	// users_map := map[string]string{"Sanya": "12341254", "Nikita": "22", "Dimon": "23"}
	// fmt.Println(reflect.DeepEqual(users, users_map))

	//1
	toolUsage := map[string]int{"Go": 3, "VSCode": 5, "Git": 2}

	for k, v := range toolUsage {
		fmt.Printf("%s использовали столько раз: %d\n", k, v)
	}

	//2
	buildStatus := map[string]bool{"Build": true, "Run": false}
	if buildStatus["Build"] == true {
		fmt.Println("Сборка прошла успешно")
	} else {
		fmt.Println("Нет сценрия видимо не успешно")
	}

	//3
	var UserName string
	fmt.Println("Введите имя")
	fmt.Scanln(&UserName)
	userInfo := map[string]interface{}{"name": UserName, "isLoggendIn": true}
	fmt.Println(userInfo)

	//4
	cpuLoad := map[int]int{1: 40, 2: 65, 3: 30}
	maxValue := 0
	for _, v := range cpuLoad {
		if maxValue < v {
			maxValue = v
		}
	}
	fmt.Printf("Самое максимум загрузки %d\n", maxValue)

	//5
	examResult := map[string]int{"name": 94, "name2": 83, "name3": 73}
	for k, v := range examResult {
		if 80 <= v {
			fmt.Printf("Сдал экзамен %s\n", k)
		} else {
			fmt.Printf("Не сдал экзамен %s\n", k)
		}
	}

	//6
	word := []string{"go", "is", "fast"}
	wordLenegth := map[string]int{}
	for _, v := range word {
		wordLenegth[v] = len(v)
	}
	fmt.Println(wordLenegth)

	//7
	var nameMenu string
	menu := map[string]int{"Burger": 2500, "Pizza": 3200, "Tea": 500}
	fmt.Println("Выберите блюдо: Burger, Pizza, Tea")
	fmt.Scan(&nameMenu)

	price, ok := menu[nameMenu]
	if !ok {
		fmt.Println("Блюдо не найдено")
	} else {
		fmt.Println("Цена", price)
	}

	//8
	loginAttempts := map[string]int{"admin": 1, "guest": 0}
	for true {
		if loginAttempts["admin"] > 2 {
			fmt.Println("Доступ заблокирован")
			break
		} else {
			fmt.Println("Попытка входа в аккаунт admin")
			loginAttempts["admin"]++
		}
	}

	//9
	sales := [2][3]int{
		{1200, 1500, 1700},
		{900, 1100, 1600},
	}
	for i := 0; i < len(sales); i++ {
		sum := 0
		for j := 0; j < len(sales[i]); j++ {
			sum += sales[i][j]
		}
		fmt.Printf("Сумма продаж магазина %d: %d\n", i+1, sum)
	}

	//10
	number1 := []int{4, 7, 2, 9, 5}
	numbersStatus2 := make(map[int]string)

	for _, v := range number1 {
		if v%2 == 0 {
			numbersStatus2[v] = "even"
		} else {
			numbersStatus2[v] = "odd"
		}
	}

	fmt.Println(numbersStatus2)
	//11
	defaultConfig := map[string]string{"host": "localhost", "port": "8080", "mode": "production"}
	currentConfig := map[string]string{"host": "localhost", "port": "8080", "mode": "production"}
	Sovp := true
	for k, v := range defaultConfig {
		if currentConfig[k] != v {
			Sovp = false
		}
	}
	if Sovp {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}

	currentConfig["mode"] = "debug"
	for k, v := range defaultConfig {
		if currentConfig[k] != v {
			Sovp = false
		}
	}
	if Sovp {
		fmt.Println("Конфигурации совпадают")
	} else {
		fmt.Println("Конфигурации отличаются")
	}
}
