package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	nowDate := time.Now()
	fmt.Println(nowDate)
	var number int = 54
	var doubbleNumber float64 = 54.5
	var str string = "Meowww"
	var isTrue bool = true
	fmt.Println(number, doubbleNumber, str, isTrue)
	meow := 10
	var a, b int = 4, 5
	fmt.Println(meow, a, b)
	fmt.Println("a + " + "b = " + strconv.Itoa(a+b))
	fmt.Println("Сумма:")
	fmt.Println(calculateSum(5.25, 6.25))
	fmt.Println("Разность:")
	fmt.Println(calculateDifference(5.25, 6.25))
	fmt.Println("Введите числа:")
	var c, d, e int
	fmt.Println("c:")
	fmt.Scan(&c)
	fmt.Println("d:")
	fmt.Scan(&d)
	fmt.Println("e:")
	fmt.Scan(&e)
	fmt.Println("Среднее число")
	fmt.Println(averageTree(c, d, e))

}
func calculateSum(a float64, b float64) float64 {
	return a + b
}
func calculateDifference(a float64, b float64) float64 {
	return a - b
}

func averageTree(a int, b int, c int) float64 {
	return (float64(a) + float64(b) + float64(c)) / 3.0
}
