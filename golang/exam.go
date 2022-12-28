package main

import (
	"fmt"
	"math"
	"sort"

	"golang.org/x/exp/slices"
)

func main() {
	MultiplicationTable(4, 6)
	MaxElem([]int{1, 4, 20, 5, 3})
	const b float64 = 2.0
	fmt.Println(Calculation(0, 3*math.Pi, 0.1*math.Pi, b))
	MaxProduct([]int{5, 0, 20, -15})
	MaxProduct2([]int{5, 0, 20, -15})
	Square(6, 4)
	OneTwo(10)
	SumElem(6)
	RockScissorsPaper(1, 3)
	fmt.Println(convertTemp(70, "C"))
	fmt.Println(MileKm(5))
	fmt.Println(KmMile(8))
	fmt.Println(Factorial(5))
	Coder([]int{1, 0, 0, 1, 1})
	Pyramid(4)
	for _, value := range TwoDimensionalAarray(4, 3) {
		fmt.Println(value)
	}
	fmt.Println(TwoDimensionalAarray(4, 3))
	const a float64 = 1.5
	fmt.Println(Calculation21(a, b, 0, 30, 2))
	fmt.Println(Odd([]int{5, 0, 20, -15}))
	fmt.Println(ArithmeticMean([]float64{5.8, 4.4, 20.5, -15.2}, 4.0))
	SortedYesOrNo([]int{90, 77, 0, -50})
	fmt.Println(AreTheySquare([]int{3, 9, 49}))
	FizzBuzz(15)
	LikesVsDislikes([]string{"Dislike", "Like", "Like", "Dislike", "Like", "Like", "Like"})
	Move10("hello")
	InsertDashes(454793)
	Bingo([10]int{8, 9, 14, 13, 6, 8, 9, 14, 13, 6})
}

func MultiplicationTable(mult, number int) { // №1
	for i := 1; i <= mult; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%v x %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func MaxElem(array []int) { // №4
	max_el := 0
	for _, value := range array {
		if max_el < value {
			max_el = value
		}
	}
	fmt.Println(max_el)
}

func Calculation(xn, xk, dx, b float64) []float64 { // №6
	var ans []float64
	for x := xn; x <= xk; x += dx {
		y := 1 + math.Sin(math.Pow(b, 3)+math.Pow(x, 3))
		ans = append(ans, y)
	}
	return ans
}

func MaxProduct(array2 []int) { // №5
	max_el := 0
	min_el := 0
	index_min := 0
	index_max := 0
	max_el1 := 0
	max_el2 := 0
	for i, value := range array2 {
		if value > max_el {
			max_el = value
			index_max = i
		}
		if min_el > value {
			min_el = value
			index_min = i
		}
	}
	maxPr1 := 0
	maxPr2 := 0
	for i, value := range array2 {
		number := value
		index_el := i
		if max_el*number > maxPr1 && index_max != index_el {
			maxPr1 = max_el * number
			max_el1 = number
		}
		if min_el*number > maxPr2 && index_min != index_el {
			maxPr2 = min_el * number
			max_el2 = number
		}
	}
	if len(array2) <= 2 {
		fmt.Println(array2)
	} else if maxPr1 > maxPr2 {
		fmt.Println(max_el, max_el1)
	} else {
		fmt.Println(min_el, max_el2)
	}
}
func MaxProduct2(array2 []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(array2)))
	max_el := array2[0]
	maxPr := array2[0] * array2[1]
	sort.Sort(sort.IntSlice(array2))
	min_el := array2[0]
	minPr := array2[0] * array2[1]
	if len(array2) <= 2 {
		fmt.Println(array2)
	} else if maxPr > minPr {
		fmt.Println(max_el, maxPr/max_el)
	} else {
		fmt.Println(min_el, minPr/min_el)
	}
}

func Square(length, width int) { // №7
	for i := 1; i <= width; i++ {
		for j := 1; j <= length; j++ {
			fmt.Print("# ")
		}
		fmt.Println()
	}
}

func OneTwo(value int) { // №9
	for i := 1; i <= value; i++ {
		fmt.Print(i, "\t")
	}
	fmt.Println()
	for i := value; i >= 1; i -= 2 {
		fmt.Print(i, "\t")
		if i > 1 {
			fmt.Print(i-1, "\t")
		}
	}
	fmt.Println()
}

func SumElem(N int) { // №10
	array2 := []int{1, 2, 3}
	for i := 0; i <= N-4; i++ {
		array2 = append(array2, array2[i]+array2[i+1]+array2[i+2])
	}
	fmt.Println(array2)
}

func RockScissorsPaper(player1, player2 int) { // №12
	if player1 == player2 && (player1 == 1 || player1 == 2 || player1 == 3) {
		fmt.Println("Ничья!")
	} else if player1 == 1 {
		if player2 == 2 {
			fmt.Println("player1 победил!")
		} else if player2 == 3 {
			fmt.Println("player2 победил!")
		}
	} else if player1 == 2 {
		if player2 == 1 {
			fmt.Println("player2 победил!")
		} else if player2 == 3 {
			fmt.Println("player1 победил!")
		}
	} else if player1 == 3 {
		if player2 == 1 {
			fmt.Println("player1 победил!")
		} else if player2 == 2 {
			fmt.Println("player2 победил!")
		}
	}
}

func convertTemp(temperature float64, convertTo string) float64 { // №14
	if convertTo == "F" {
		temperature = 9.0/5.0*temperature + 32
	} else if convertTo == "C" {
		temperature = 5.0 / 9.0 * (temperature - 32)
	}
	return temperature
}

func MileKm(length float64) float64 { // №15
	length = length * 1.609
	return length
}
func KmMile(length float64) float64 {
	length = length / 1.609
	return length
}

func Factorial(number int) int { // №16
	fact := 1
	for i := 2; i <= number; i++ {
		fact = fact * i
	}
	return fact
}

func Coder(array []int) { // №17
	var newCode []int
	for i, value := range array {
		if i == 0 {
			if value == 0 {
				newCode = append(newCode, 0)
			} else {
				newCode = append(newCode, 1)
			}
		} else {
			if array[i-1] == value {
				newCode = append(newCode, 0)
			} else {
				newCode = append(newCode, 1)
			}
		}
	}
	fmt.Println(newCode)
}

func Pyramid(heigth int) { // №18
	for i := 0; i <= heigth-1; i++ {
		for j := 1; j <= 2*heigth-1; j++ {
			if j == heigth-i {
				for k := 0; k <= i*2; k++ {
					fmt.Print("#")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func TwoDimensionalAarray(n, m int) [][]int { // №19
	var tArray [][]int
	var arr []int
	for i := 0; i <= n-1; i++ {
		arr = []int{}
		for j := i*m + 1; j <= m*(i+1); j++ {
			arr = append(arr, j)
		}
		tArray = append(tArray, arr)
	}
	return tArray
}

func Calculation21(a, b, xn, xk, n float64) []float64 { // №20
	var ans []float64
	for x := xn; x <= xk; x += n {
		y := math.Pow((a*x + b), 1.0/3.0)
		ans = append(ans, y)
	}
	return ans
}

func Odd(array []int) []int { // №21
	var oddNum []int
	for _, value := range array {
		if value%2 != 0 {
			oddNum = append(oddNum, value)
		}
	}
	return oddNum
}

func ArithmeticMean(array []float64, size float64) float64 { // №22
	sum := 0.0
	for _, value := range array {
		sum += value
	}
	return sum / size
}

func SortedYesOrNo(array []int) { //codewars ||
	count := 0 //         \/
	for i := 1; i < len(array); i++ {
		if array[i] <= array[i-1] {
			count += -1
		} else if array[i] >= array[i-1] {
			count += 1
		}
	}
	if count == len(array)-1 {
		fmt.Println("Yes, ascending")
	} else if count == 0-len(array)+1 {
		fmt.Println("Yes, descending")
	} else {
		fmt.Println("No")
	}
}

func AreTheySquare(array []int) string {
	var ans string
	if len(array) == 0 {
		ans = "undefined"
	} else {
		for _, value := range array {
			if math.Mod(float64(value), math.Sqrt(float64(value))) != 0 {
				ans = "false"
				break
			} else {
				ans = "true"
			}
		}
	}
	return ans
}

func FizzBuzz(N int) {
	var arr []string
	for i := 1; i <= N; i++ {
		if i%3 == 0 && i%5 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, fmt.Sprintf("%v", i))
		}
	}
	fmt.Println(arr)
}

func LikesVsDislikes(array []string) {
	var result string = "Nothing"
	for i := 1; i < len(array); i++ {
		if array[i] == array[i-1] {
			if result == "Nothing" {
				result = array[i]
			} else {
				result = "Nothing"
			}
		} else {
			result = array[i]
		}
	}
	fmt.Println(result)
}

func Move10(message string) {
	for i := 0; i < len(message); i++ {
		word := message[i]
		if word > 'a' && word < 'z' {
			word = word + 10
			if word > 'z' {
				word = word - 26
			}
		}
		fmt.Printf("%c", word)
	}
	fmt.Println()
}

func InsertDashes(num int) {
	number := fmt.Sprint(num)
	for i := 1; i < len(number); i++ {
		if number[i]%2 != 0 && number[i-1]%2 != 0 {
			fmt.Printf(fmt.Sprintf("-%c", number[i]))
		} else {
			if i == 1 {
				fmt.Print(fmt.Sprintf("%c%c", number[i-1], number[i]))
			} else {
				fmt.Print(fmt.Sprintf("%c", number[i]))
			}
		}
	}
	fmt.Println()
}

func Bingo(array [10]int) {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var win []string
	for _, value := range array {
		if alphabet[value] == 'B' {
			win = append(win, fmt.Sprint("B"))
		}
		if alphabet[value] == 'I' {
			win = append(win, fmt.Sprint("I"))
		}
		if alphabet[value] == 'N' {
			win = append(win, fmt.Sprint("N"))
		}
		if alphabet[value] == 'G' {
			win = append(win, fmt.Sprint("G"))
		}
		if alphabet[value] == 'O' {
			win = append(win, fmt.Sprint("O"))
		}
	}
	if slices.Contains(win, "B") && slices.Contains(win, "I") && slices.Contains(win, "N") && slices.Contains(win, "G") && slices.Contains(win, "O") {
		fmt.Println("WIN")
	} else {
		fmt.Println("LOSE")
	}
}
