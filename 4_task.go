package main

import (
	"fmt"
	"strconv"
)

type MultTable struct {
	size int
	matrix [][]int
}

func InitTable(n int) *MultTable {
	newMatrix := make([][]int, n + 1)
	newMatrix[0] = make([]int, n + 1)
	for i := 0; i <= n; i++ {
		newMatrix[0][i] = i
	}
	for i := 1; i <= n; i++ {
		newMatrix[i] = make([]int, n + 1)
		newMatrix[i][0] = i
		for e := 1; e <= n; e++ {
			newMatrix[i][e] = i*e
		}
	}
	return &MultTable{size: n, matrix: newMatrix}
}

func (mT *MultTable) PrintTeble() {
	maxLen := make([]int, mT.size + 1)
	for i := 0; i <= mT.size; i++ {
		// чтобы подсчитать кол-во цифр в числе
		num := mT.matrix[mT.size][i]
		strNum := strconv.Itoa(num)
		maxLen[i] = len(strNum)
	}
	for i := 0; i <= mT.size; i++ {
		for e := 0; e <= mT.size; e++ {
			num := mT.matrix[i][e]
			strNum := strconv.Itoa(num)
			diff := maxLen[e] - len(strNum)
			for w := 0; w < diff; w++ {
				fmt.Print(" ")
			}
			if num == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(strNum)
			}
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	table := InitTable(n)
	table.PrintTeble()
}
