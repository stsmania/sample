package utility

import (
	"fmt"
	"strconv"
)

func ConvertStringSliceToIntSlice(numericStringList []string) ([]int, error) {
	var numbers []int
	for _, str := range numericStringList {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("エラー: %v\n", err)
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}
