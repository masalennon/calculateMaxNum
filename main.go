package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const maxInt = math.MaxInt64
const minInt = math.MinInt64

func main() {
	inputNum := readStdn()
	fmt.Println(calculateResult(inputNum))
}

func calculateResult(input string) int {
	slice := strings.Fields(strings.ReplaceAll(input, ",", " "))
	if len(slice) < 3 {
		os.Exit(0)
	}
	numSlice := convertSlice(slice)
	return calculateMaxComb(numSlice)
}

func convertSlice(s []string) []int {
	ints := make([]int, len(s))
	for i, s := range s {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func readStdn() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func calculateMaxComb(slice []int) int {
	maxs := max(slice)
	mins := min(slice)
	if maxs[0] < 0 && maxs[1] < 0 && maxs[2] < 0 { //全て負の数の場合
		return maxs[0] * maxs[1] * maxs[2]
	}
	if mins[0]*mins[1] > maxs[1]*maxs[2] { //1, 2番目に小さい数の掛け算が、2, 3番目に大きい数の掛け算よりも大きい時
		return maxs[0] * mins[0] * mins[1]
	}
	return maxs[0] * maxs[1] * maxs[2]
}

func min(n []int) [3]int { //-11 -12 -33 -111 -111 -33
	min := maxInt
	var secondMin, thirdMin int
	for _, i := range n {
		if i <= min {
			thirdMin = secondMin
			secondMin = min
			min = i
		}
		if i <= secondMin && i > min {
			thirdMin = secondMin
			secondMin = i
		}
		if i <= thirdMin && i > secondMin && i > min {
			thirdMin = i
		}
	}
	threeMins := [3]int{min, secondMin, thirdMin}
	return threeMins
}

func max(n []int) [3]int {
	max := int(minInt)
	var secondMax, thirdMax int
	for _, i := range n {
		if i >= max {
			thirdMax = secondMax
			secondMax = max
			max = i
		}
		if i >= secondMax && i < max {
			thirdMax = secondMax
			secondMax = i
		}
		if i >= thirdMax && i < secondMax && i < max {
			thirdMax = i
		}
	}
	threeMaxs := [3]int{max, secondMax, thirdMax}
	return threeMaxs
}
