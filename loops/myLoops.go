package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// for i := range days {
	// 	fmt.Println(days[i])

	// 	for _, day := range days {
	// 		fmt.Println("Index is %v and value is %v\n", index, day)
	// 	}
	// }

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			break

		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}

}
