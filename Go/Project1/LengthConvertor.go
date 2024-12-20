package main

import (
	"fmt"
	"strconv"
	"strings"
)

var UNIT_VALUES = map[string]float32{"m": 1, "cm": 0.01, "mm": 0.001, "km": 1000}
var UNIT_NAMES = map[string]string{"m": "metre", "cm": "centimetre", "mm": "milimetre", "km": "kilometre"}

func changeUnit(text string, originalUnit string) string {
	var input string

	for {
		fmt.Println("Input a unit to change to for", text)
		fmt.Scan(&input)
		if inMap(input) {
			break
		} else {
			fmt.Println("Invalid input, please use a valid unit")
		}
	}

	fmt.Printf("Unit changed from %v (%v) to %v (%v)\n", originalUnit, UNIT_NAMES[originalUnit], input, UNIT_NAMES[input])
	return input
}

func inMap(input string) bool {
	_, ok := UNIT_NAMES[input]
	return ok
}

func checkSame(currentUnit string, convertUnit string) bool {
	return currentUnit != convertUnit
}

func checkFloat(input string) bool {
	_, err := strconv.ParseFloat(input, 10)

	return err != nil || input == ""
}

func testMap() {
	var input string

	fmt.Scan(&input)
	var state bool = inMap(input)
	fmt.Println(state)
}

func main() {

	var currentUnit string = "m"
	var convertUnit string = "cm"
	var menuChoice string
	const CURRENTUNITTEXT string = "Current unit"
	const CONVERTUNITTEXT string = "Convert unit"

	for {
		fmt.Println("Welcome to length convertor")
		fmt.Printf("Current Mode: %v(%v) to %v(%v) \n", currentUnit, UNIT_NAMES[currentUnit], convertUnit, UNIT_NAMES[convertUnit])

		if !checkSame(currentUnit, convertUnit) {
			fmt.Println("WARNING: Your current unit and convert unit are the same, please adjust your convert mode")
		}

		fmt.Println("Select one of the following options:")
		fmt.Println("1. Convert value")
		fmt.Println("2. Change unit")
		fmt.Println("3. Quit")
		fmt.Print("Your Input: ")
		fmt.Scan(&menuChoice)
		switch menuChoice {
		case "1":
			length_convert(currentUnit, convertUnit)

		case "2":
			currentUnit = changeUnit(CURRENTUNITTEXT, currentUnit)
			convertUnit = changeUnit(CONVERTUNITTEXT, convertUnit)
			fmt.Println("Your convert mode have been changed.")
		case "3":
			fmt.Println("Goodbye")

		default:
			fmt.Println("Invalid choice")
		}

		if menuChoice == "3" {
			break
		}
	}
}

func length_convert(currentUnit string, convertUnit string) {
	var length_value string

	for {
		fmt.Println("Please input your length")
		fmt.Print("Your input:")
		fmt.Scan(&length_value)
		if checkFloat(length_value) {

		} else {
			length, _ := strconv.ParseFloat(strings.TrimSpace(length_value), 10)
			output := float32(length) * float32(UNIT_VALUES[currentUnit]) / float32(UNIT_VALUES[convertUnit])
			fmt.Printf("%v %v (%v) \n", output, UNIT_NAMES[convertUnit], convertUnit)
			break
		}

	}
}
