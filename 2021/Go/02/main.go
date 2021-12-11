package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type moveOrder struct {
	direction string
	distance  int
}

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

	navigator := navigate()

	total, navErr := navigator(inputFile)
	if navErr != nil {
		fmt.Errorf("unable to calculate: %s\n", navErr.Error())
	}

	fmt.Printf("the calculated total is: %d\n", total)
}

func calculate() func(order *moveOrder, fb, ud int) (h, v int) {
	aim := 0

	return func(order *moveOrder, fb, ud int) (h, v int) {
		switch order.direction {
		case "forward":
			return fb + order.distance, ud + (aim * order.distance)
		case "up":
			aim = aim - order.distance
			return fb, ud
		case "down":
			aim = aim + order.distance
			return fb, ud
		}

		return fb, ud
	}
}

func navigate() func(orders *os.File) (int, error) {
	h, v := 0, 0

	return func(orders *os.File) (int, error) {
		oom := orderOfMarch(orders)
		calculator := calculate()

		for _, order := range *oom {
			mo := new(moveOrder)
			move := strings.Split(order, " ")

			distance, convErr := strconv.Atoi(move[1])
			if convErr != nil {
				return 0, fmt.Errorf("unable to convert to integer: %v\n", convErr.Error())
			}

			mo.direction = move[0]
			mo.distance = distance
			h, v = calculator(mo, h, v)
		}

		return h * v, nil
	}
}

func orderOfMarch(orders *os.File) *[]string {
	var orderSeries []string

	scanner := bufio.NewScanner(orders)
	for scanner.Scan() {
		orderSeries = append(orderSeries, scanner.Text())
	}

	return &orderSeries
}
