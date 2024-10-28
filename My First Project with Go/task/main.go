package main

import "fmt"

func main() {
	itemsEarnings := map[string]float64{
		"Bubblegum":      202,
		"Toffee":         118,
		"Ice cream":      2250,
		"Milk chocolate": 1680,
		"Doughnut":       1075,
		"Pancake":        80,
	}
	var totalEarned, staffExpenses, otherExpenses, netIncome float64 = 0, 0, 0, 0

	fmt.Println("Earned amount:")
	for _, itemEarned := range itemsEarnings {
		totalEarned += itemEarned
	}
	fmt.Printf("Bubblegum: $%.f\n", itemsEarnings["Bubblegum"])
	fmt.Printf("Toffee: $%.f\n", itemsEarnings["Toffee"])
	fmt.Printf("Ice cream: $%.f\n", itemsEarnings["Ice cream"])
	fmt.Printf("Milk chocolate: $%.f\n", itemsEarnings["Milk chocolate"])
	fmt.Printf("Doughnut: $%.f\n", itemsEarnings["Doughnut"])
	fmt.Printf("Pancake: $%.f\n", itemsEarnings["Pancake"])
	fmt.Printf("\nIncome: $%.1f\n", totalEarned)
	fmt.Println("Staff expenses:")
	fmt.Scan(&staffExpenses)
	fmt.Println("Other expenses:")
	fmt.Scan(&otherExpenses)
	netIncome = totalEarned - (staffExpenses + otherExpenses)
	fmt.Printf("Net income: %.f", netIncome)
}
