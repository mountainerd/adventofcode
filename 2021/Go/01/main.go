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

	defer func() {
		err := inputFile.Close()
		if err != nil {
			fmt.Printf("closing did not go according to plan: %s\n", err.Error())
			os.Exit(1)
		}
	}()

	readings, err := convertStringsToIntegers(inputFile)
	if err != nil {
		fmt.Printf("unable to complete conversion:\n\t%s\n", err.Error())
		os.Exit(1)
	}

	windowReadings, err := windowTotals(readings)
	if err != nil {
		fmt.Printf("unable to complete conversion:\n\t%s\n", err.Error())
	}

	fmt.Printf("total count: %d\n", depthCounter(windowReadings))
}

func adder(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
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

func windowTotals(measurements *[]int) (*[]int, error) {
	var windowMeasurements []int

	for i := 3; i <= len(*measurements); i++ {
		windowMeasurements = append(windowMeasurements, adder((*measurements)[i-3:i]))
	}

	return &windowMeasurements, nil
}
