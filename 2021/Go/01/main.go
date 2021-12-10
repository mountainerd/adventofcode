package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("unable to open inputFile: %s\n", err.Error())
		os.Exit(1)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("close already called: %s\n", err.Error())
			os.Exit(1)
		}
	}(inputFile)

	readings, err := convertStringsToIntegers(inputFile)
	if err != nil {
		fmt.Printf("unable to complete conversion:\n\t%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("total count: %d\n", depthCounter(readings))
}

func convertStringsToIntegers(f *os.File) (*[]int, error) {
	var measurements []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, errors.New(err.Error())
		}
		measurements = append(measurements, i)
	}

	return &measurements, nil
}

func depthCounter(depthList *[]int) int {
	var depthCount int

	depthChecker := isDeeper()

	for _, depth := range *depthList {
		depthCount = depthChecker(depth)
	}

	return depthCount
}

func isDeeper() func(depth int) int {
	lastMeasurement, isLarger := 0, 0

	return func(depth int) int {
		if depth > lastMeasurement {
			isLarger += 1
		}

		lastMeasurement = depth
		return isLarger - 1
	}
}
