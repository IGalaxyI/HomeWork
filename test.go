package main

import "fmt"

type Devil struct {
	Name     string
	Weight   float64
	Distance float64
	Count_H  float64
}

func main() {
	const BaseRate = 5.50 //за 1 кг
	const TaxRate = 0.12
	const Distancerate = 2.0 //за 1 км
	const FragileFee = 0.2

	var devil Devil

	fmt.Println("Имя:")
	fmt.Scan(&devil.Name)
	fmt.Println("Вес:")
	fmt.Scan(&devil.Weight)
	fmt.Println("Дистанция:")
	fmt.Scan(&devil.Distance)
	fmt.Println("Хрупкого кол-во:")
	fmt.Scan(&devil.Count_H)

	Sum := (devil.Weight*BaseRate)*(1+FragileFee*devil.Count_H) + (devil.Distance * Distancerate)
	SumTax := Sum + (Sum * TaxRate)

	fmt.Println(SumTax)

}
