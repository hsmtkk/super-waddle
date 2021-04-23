package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Enter your weight (kg): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	weight, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Weight on Earth: %f kg\n", weight)

	marsWeight := calculateWeightOnMars(weight)
	fmt.Printf("Weight on Mars: %f kg\n", marsWeight)
}

func calculateWeightOnMars(weight float64) float64 {
	return weight / 9.81 * 3.711
}
