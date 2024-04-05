package main

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножествах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	const step = 10

	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int64][]float64)

	minTemperature := slices.Min(array)
	maxTemperature := slices.Max(array)

	minKey := getTemperaturePoint(minTemperature)
	maxKey := getTemperaturePoint(maxTemperature)

	for currentKey := minKey; currentKey <= maxKey; currentKey += step {
		var tempSlice []float64

		currentKeyFloat := float64(currentKey)
		nextKey := float64(currentKey + step)

		for _, value := range array {
			if (value > currentKeyFloat || isFloatEqual(value, currentKeyFloat)) && value < nextKey {
				tempSlice = append(tempSlice, value)
			}
		}

		if len(tempSlice) > 0 {
			groups[currentKey] = tempSlice
		}
	}

	str, err := json.MarshalIndent(groups, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("\n", string(str))
}

func getTemperaturePoint(value float64) int64 {
	var temperaturePoint int64

	if value < 0 {
		temperaturePoint = int64(value/10)*10 - 10
	} else {
		temperaturePoint = int64(value/10) * 10
	}

	return temperaturePoint
}

func isFloatEqual(a float64, b float64) bool {
	const epsilon = 0.00000001
	return math.Abs((a-b)/b) < epsilon
}

/*
 {
  "-30": [
    -25.4,
    -27,
    -21
  ],
  "10": [
    13,
    19,
    15.5
  ],
  "20": [
    24.5
  ],
  "30": [
    32.5
  ]
}
*/
