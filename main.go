package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog","d","dor"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i,j}
			}
		}
	}

	return nil
}


func reverse(x int) int {
	var result int64 = 0

	for x != 0 {
		result = result * int64(10) + int64(x % 10)
		x /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	} else {
		return int(result)
	}
}

func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) { return false}
	var reverse int = 0
	for x > reverse{
		reverse = reverse * 10 + x % 10
		x /= 10
	}
	return x == reverse || x == reverse / 10
}

func romanToInt(s string) int {
	var result int = 0
	length := len(s)
	for i := 0; i < length; i++ {
		c := s[i]
		var next uint8 = ' '
		if i+1 < length {
			next = s[i+1]
		}
		switch c {
		case 'I':
			if (next == 'V' || next == 'X') {
				result -= 1;
			} else {
				result += 1;
			}
		case 'V':
			result += 5;
		case 'X':
			if ( next == 'L' || next == 'C') {
				result -= 10;
			} else {
				result += 10;
			}
		case 'L':
			result += 50;
		case 'C':
			if ( next == 'D' || next == 'M') {
				result -= 100;
			} else {
				result += 100;
			}
		case 'D':
			result += 500;
		case 'M':
			result += 1000;
		}
	}
	return result
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 { return ""}
	first := strs[0]
	var i int = 0
	for ;i<len(first);i++ {
		for j:=1;j<len(strs);j++ {
			if len(strs[j]) <= i || first[i] != strs[j][i] {
				return first[0:i]
			}
		}
	}
	return first[0:i]
}