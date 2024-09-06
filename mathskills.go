package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var Population []string

func main() {
	content, err := os.ReadFile("popdata.txt")
	if err != nil {
		log.Fatal("couldn't read file")
	}
	Population = strings.Split(string(content), "\n")
	fmt.Println("average is ", Average())

	fmt.Println("//////////////////////////////////////////")
	fmt.Println()

	ordered, mediane := Mediane()
	fmt.Println("ordered is : ", ordered)
	fmt.Println()

	fmt.Println("//////////////////////////////////////////")
	fmt.Println()

	fmt.Println("mediane is : ", mediane)

	fmt.Println("//////////////////////////////////////////")
	fmt.Println()

	variance := Variance(Population)
	fmt.Println("Variance is : ", variance)

	fmt.Println("//////////////////////////////////////////")
	fmt.Println()

	deviation := StandardDeviation()
	fmt.Println("Standard Deviation is : ", deviation, "len pop is :", len(Population))
	fmt.Println("\nlen/2-1", ordered[(len(ordered)/2-1)], "\nlen/2", ordered[(len(ordered)/2)], "\n101", Population[101], "100", Population[100])
	st := "12345"
	fmt.Print(Atoi(st))
}

func Atoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var s string
	var sign int = 1
	for i := 0; i < len(str); i++ {
		if string(str[0]) == "-" {
			sign = -1
		}
		if string(str[i]) >= "0" && string(str[i]) <= "9" {
			s += string(str[i])
		}
	}
	var c int
	var n int
	n = int(s[0]-48) * 10
	for i := 0; i < len(s)-1; i++ {

		c = n + int(s[i+1]-48)
		n = c * 10

	}

	return (n / 10) * sign
}

func Average() int {
	var sum int
	var average int
	for _, value := range Population {
		sum += Atoi(value)
	}
	average = sum / len(Population)
	return average
}

func Mediane() ([]int, int) {
	var ordered []int
	var mediane int
	for i := 0; i < len(Population)-1; i++ {
		for j := 0; j < len(Population)-1; j++ {
			if Atoi(Population[j]) > Atoi(Population[j+1]) {
				Population[j], Population[j+1] = Population[j+1], Population[j]
			}
		}
	}

	for _, value := range Population {
		ordered = append(ordered, Atoi(value))
	}

	if len(ordered)%2 == 0 {
		mediane = (ordered[(len(ordered)/2-1)] + ordered[(len(ordered)/2)]) / 2
	} else {
		mediane = (ordered[len(ordered)/2])
	}
	return ordered, mediane
}

func Variance([]string) int {
	var res int
	var variance int
	average := Average()

	for _, value := range Population {
		res = (Atoi(value) - average) * (Atoi(value) - average)
		variance += res
	}
	variance = variance / len(Population)
	return variance
}

func StandardDeviation() int {
	var deviation int
	variance := Variance(Population)
	deviation = int(math.Sqrt(float64(variance)))
	return deviation
}

// variance := Variance(Population)
// var deviation float64
// var i float64 = 1
// for i < float64(variance) {
// 	if i*i == float64(variance) {
// 		deviation = i
// 	}
// 	i++
// }
// return deviation
