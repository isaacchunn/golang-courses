package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) float64 {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		//We can just return a default value of 1000
		return defaultValue
	}
	valueText := string(bs)
	value, _ := strconv.ParseFloat(valueText, 64)
	return value
}
