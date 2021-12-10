package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		errors.New(err.Error())
	}
	defer file.Close()

	var measurements []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		measurements = append(measurements, i)
	}

	for _, measurement := range measurements {
		fmt.Println(isDeeper(measurement))
	}
}

func isDeeper(depth int) func() int {
	return func() int {
		//lastMeasurement, isLarger := 0, 0
		//
		//if depth > lastMeasurement {
		//	isLarger += 1
		//}
		//
		//lastMeasurement = depth
		//return isLarger - 1
		return depth
	}
}
