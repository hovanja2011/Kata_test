package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func oper(a int, b int, o string) int {
	switch o {
	case "+":
		return (a + b)
	case "-":
		return (a - b)
	case "*":
		return (a * b)
	case "/":
		return (a / b)
	default:
		return (-100)
	}
}

var rimTensMap = map[int]string{10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C"}
var rimArr = [11]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "IIX", "IX", "X"}
var rimMap = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

func similiarTest(a, b string) bool {
	_, err_a := strconv.Atoi(a)
	_, err_b := strconv.Atoi(b)
	switch {
	case (err_a == nil) && (err_b == nil):
		return true
	case (err_a != nil) && (err_b != nil):
		return false
	}
	err := errors.New("error: одновременно используются разные системы счисления")
	fmt.Print(err)
	os.Exit(0)
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		expression, _ := reader.ReadString('\n')
		expression = strings.Replace(expression, "\n", "", -1)
		expr_arr := strings.Split(expression, " ")
		if len(expr_arr) != 3 {
			err := errors.New("error: формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
			fmt.Print(err)
			os.Exit(0)
		}
		var result = ""
		o := expr_arr[1]
		switch similiarTest(expr_arr[0], expr_arr[2]) {
		case true:
			{
				a_num, _ := strconv.Atoi(expr_arr[0])
				b_num, _ := strconv.Atoi(expr_arr[2])
				if (a_num > 10) || (b_num > 10) {
					err := errors.New("error: значение операндов более 10")
					fmt.Print(err)
					os.Exit(0)
				}
				if (a_num < 1) || (b_num < 1) {
					err := errors.New("error: значение операндов менее 1")
					fmt.Print(err)
					os.Exit(0)
				}
				result = strconv.Itoa(oper(a_num, b_num, o))
				fmt.Println(result)
			}
		case false:
			{
				a_num, err_a := rimMap[expr_arr[0]]
				b_num, err_b := rimMap[expr_arr[2]]
				if (err_a == true) || (err_b == true) {
					err := errors.New("error: латинские операнды записаны некорректно, ожидаются числа I .. X ")
					fmt.Print(err)
					os.Exit(0)
				}
				c := oper(a_num, b_num, o)
				var d, m int = 0, 0
				if c < 1 {
					err := errors.New("error: в римской системе нет отрицательных чисел")
					fmt.Print(err)
					os.Exit(0)
				}
				if c <= 10 {
					result = rimArr[c]
					fmt.Println(result)
				}
				if c > 10 {
					d = c / 10
					m = c % 10
					result = rimTensMap[d*10] + rimArr[m]
					fmt.Println(result)
				}
			}
		}
	}
}
