package day04

import (
	"fmt"
)

func findContinuousSequence(target int) [][]int {
	res := [][]int{}

	for i := 1; i <= (target+1)/2; i++ {
		for j := target/i - 1; j <= target; j++ {
			if (2*i+j-1)*j/2 == target {
				res = append(res, []int{i})
				for k := 1; k < j; k++ {
					res[len(res)-1] = append(res[len(res)-1], i+k)
				}
			}
			if (2*i+j-1)*j/2 > target {
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findContinuousSequence(20))
}
