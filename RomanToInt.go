package main

import ()

func romanToInt(s string) int {
	// Mapping of Roman numerals to their integer values
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prevValue := 0

	// Iterate over the string from right to left
	for i := len(s) - 1; i >= 0; i-- {
		value := romanMap[s[i]]

		// If the current value is less than the previous value, subtract it from the total
		// Otherwise, add it to the total
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total
}

// func RomanToInt() {
// 	// Test cases
// 	romanToInttestCases := []string{"III", "LVIII", "MCMXCIV"}

// 	for _, testCase := range romanToInttestCases {
// 		result := romanToInt(testCase)
// 		fmt.Printf("%s = %d\n", testCase, result)
// 	}
// }
