package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	var diagnosticReports []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		diagnosticReports = append(diagnosticReports, scanner.Text())
	}

	frequencies := frequencyCheck(&diagnosticReports)

	fmt.Println("total power consumption:", calculatePowerConsumption(&frequencies))
}

func calculatePowerConsumption(frequencies *[]int) int {
	var eRate string
	var gRate string

	for _, reading := range *frequencies {
		if reading <= 0 {
			eRate += "0"
			gRate += "1"
		} else {
			eRate += "1"
			gRate += "0"
		}
	}

	epsilon, _ := strconv.ParseInt(eRate, 2, 32)
	gamma, _ := strconv.ParseInt(gRate, 2, 32)
	return int(epsilon * gamma)
}

func frequencyCheck(readings *[]string) []int {
	binaryTracker := make([]int, 12)

	for _, binaryReading := range *readings {
		componentReading := strings.Split(binaryReading, "")

		for index, value := range componentReading {
			if value == "0" {
				binaryTracker[index] -= 1
			} else {
				binaryTracker[index] += 1
			}
		}
	}

	return binaryTracker
}

//func findOxygenGeneratorRating(frequencies *[]int)
