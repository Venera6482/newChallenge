package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices lesson")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Type of fruitless is %T\n", fruitList)
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	highScores := make([]int, 4)
	highScores[0] = 345
	highScores[1] = 347
	highScores[2] = 754
	highScores[3] = 932                            //highScores[4]=8844
	highScores = append(highScores, 555, 666, 777) //can
	fmt.Println(highScores)
	sort.Ints(highScores) //available only in sorts    fmt.Println(highScores)
	fmt.Println(highScores)
}
