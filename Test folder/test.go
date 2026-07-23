package main

import "fmt"

func main() {
	smartHome()
}

type Product struct {
	Name  string
	Power float64
	Time  float64
	Night bool
}

func smartHome() {
	const BaseTariff = 0.45
	const HighLoadTax = 0.15
	const NightDiscount = 0.30

	var products []Product

	for true {
		var product Product
		fmt.Println("Введите название прибора:")
		fmt.Scan(&product.Name)
		if product.Name == ("done") {
			break
		}
		fmt.Println("Введите мощьность:")
		fmt.Scan(&product.Power)
		fmt.Println("Введите время:")
		fmt.Scan(&product.Time)
		fmt.Println("Введите ночной режим true/false:")
		fmt.Scan(&product.Night)
		products = append(products, product)
	}

	for _, product := range products {
		Rashod := (product.Power * product.Time) / 1000
		Cash := Rashod * BaseTariff
		category := ""
		if product.Night {
			Cash = Cash - (Cash * NightDiscount)
		}
		if Rashod > 10 {
			Cash = Cash + (Cash * HighLoadTax)
		}
		switch {
		case product.Power >= 0 && product.Power < 100:
			category += "Экономный"
		case product.Power > 100 && product.Power < 1000:
			category += "Стандартный"
		case product.Power > 1000:
			category += "Мощный"
		}

		fmt.Printf("--Отчет по прибору--\nПрибор: %s [Категория:%s] \nРасход: %.2f \nИтоговая стоимость: %.2f\n\n", product.Name, category, Rashod, Cash)
	}
}
