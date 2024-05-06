package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{-9, -18, 45, 27, 18}
	fmt.Println(commonDiv(numbers))
}

func commonDiv(nums []int) []int {
	ans := []int{}
	num := nums[0]
	// заполняем только для первого числа
	// я слелал, чтобы работало для отрицательных чисел,
	// но делители были только положительными
	for e := 1; e <= int(math.Pow(math.Abs(float64(num)), 0.5)); e++ {
		// если сразу начинать с двух, из-за корня
		// цикл будет срабатывать только с 4
		if e == 1 { continue }
		if num % e == 0 {
			ans = append(ans, e)
			secondDiv := int(math.Abs(float64(num/e)))
			// проверка не сработает в случае, когда число - квадрат
			if secondDiv != e {
				ans = append(ans, secondDiv)
			}
		}
	}
	// само число тоже может быть делителем
	if num != 1 {
		ans = append(ans, int(math.Abs(float64(num))))
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
