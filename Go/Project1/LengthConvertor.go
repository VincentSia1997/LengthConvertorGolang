package main

import (
	"fmt"
)

var UNIT_VALUES = map[string]float32{"metre": 1, "centimetre": 0.01, "milimetre": 0.001, "kilometre": 1000}
var UNIT_NAMES = map[string]string{"m": "metre", "cm": "centimetre", "mm": "milimetre", "km": "kilometre"}

func changeUnit(originalUnit string) string {
	var input string

	for {
		fmt.Println("Input a unit to change to")
		fmt.Scan(&input)
		if inMap(input) {
			break
		} else {
			fmt.Println("Invalid input, please use a valid unit")
		}
	}

	fmt.Printf("Unit changed from %v to %v\n", originalUnit, input)
	return input
}

func inMap(input string) bool {
	_, ok := UNIT_NAMES[input]
	return ok
}

func checkSame(currentUnit string, convertUnit string) bool {
	return currentUnit == convertUnit
}

func testMap() {
	var input string

	fmt.Scan(&input)
	var state bool = inMap(input)
	fmt.Println(state)
}

func main() {

	// var test float32 = 10
	var length_value float32
	var output float32
	var currentUnit string = "m"
	var convertUnit string = "cm"
	var choice string

	for {
		fmt.Println("Welcome to length convertor")
		fmt.Printf("Current Mode: %v(%v) to %v(%v) \n", currentUnit, UNIT_NAMES[currentUnit], convertUnit, UNIT_NAMES[convertUnit])

		if checkSame(currentUnit, convertUnit) == true {
			fmt.Println("WARNING: Your current unit and convert unit are the same, please adjust your convert mode")
		}

		fmt.Println("Select one of the following options:")
		fmt.Println("1. Convert value")
		fmt.Println("2. Change unit")
		fmt.Println("3. Quit")

		fmt.Scan(&choice)
		switch choice {
		case "1":
			// test = test * UNIT_VALUES[UNITS_NAMES[currentUnit]] / UNIT_VALUES[UNITS_NAMES[convertUnit]]
			fmt.Println("Please input your length")
			fmt.Scan(&length_value)
			output = length_value * UNIT_VALUES[currentUnit] / UNIT_VALUES[convertUnit]
			fmt.Printf("%#v %v (%v) \n", output, UNIT_NAMES[convertUnit], convertUnit)

		case "2":
			currentUnit = changeUnit(currentUnit)
			convertUnit = changeUnit(convertUnit)
			fmt.Println("Your convert mode have been changed.")
		case "3":
			fmt.Println("Goodbye")

		default:
			fmt.Println("Invalid choice")
		}

		if choice == "3" {
			break
		}
	}
}
