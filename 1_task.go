package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{-18, 42, 12, 18}
	fmt.Println(commonDiv(numbers))
}

func commonDiv(nums []int) []int {
	ans := []int{}
	num := nums[0]
	// заполняем только для первого числа
	// я слелал, чтобы работало для отрицательных чисел,
	// но делители были только положительными
	for e := 2; e <= int(math.Pow(math.Abs(float64(num)), 0.5)); e++ {
		if num % e == 0 {
			ans = append(ans, e)
			ans = append(ans, int(math.Abs(float64(num/e))))
		}
	}
	for _, num := range nums {
		// идём с конца чтобы удаляя не ломать индексы
		for i := len(ans)-1; i >= 0; i-- {
			// если делитель одного числа не делитель другого - удаляем
			if num % ans[i] != 0 {
				ans = append(ans[:i], ans[i+1:]...)
			}
		}
	}
	return ans
}
