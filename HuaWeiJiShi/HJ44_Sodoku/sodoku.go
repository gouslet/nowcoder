/*
 * File: \sodoku.go                                                            *
 * Project: newcoder                                                           *
 * Created At: Tuesday, 2022/04/19 , 20:55:15                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/04/20 , 02:08:07                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"fmt"
	"reflect"
)

func Sodoku(nums [9][9]uint) [9][9]uint {
	res := [9][9]uint{}

	var soduku func(nums [9][9]uint, i, j int)
	soduku = func(nums [9][9]uint, i, j int) {
		if zeroCount(nums) == 0 {
			res = nums
			return
		}
		if i > 8 && j > 8 {
			return
		}
		// printMatrix(nums)
		// fmt.Println("---------------------------")
		if n := nums[i][j]; n == 0 {
			inter := intersection(notInRow(nums, i), notInColumn(nums, j), notInSub(nums, i, j))
			if l := len(inter); l == 0 {
				return
			} else if l == 1 {
				nums[i][j] = inter[0]
				if i == 8 {
					if j == 8 {
						res = nums
						return
					} else {
						soduku(nums, i, j+1)
					}
				} else {
					if j == 8 {
						soduku(nums, i+1, 0)
					} else {
						soduku(nums, i, j+1)
					}
				}

			} else {
				for _, y := range inter {
					nums[i][j] = y
					if i == 8 {
						if j == 8 {
							res = nums
							return
						} else {
							soduku(nums, i, j+1)
						}
					} else {
						if j == 8 {
							soduku(nums, i+1, 0)
						} else {
							soduku(nums, i, j+1)
						}
					}

				}
			}
		} else {
			if i == 8 {
				if j == 8 {
					res = nums
					return
				} else {
					soduku(nums, i, j+1)
				}
			} else {
				if j == 8 {
					soduku(nums, i+1, 0)
				} else {
					soduku(nums, i, j+1)
				}
			}

		}
	}
	// for i, row := range nums {
	// 	for j, n := range row {

	// 	}

	// }
	soduku(nums, 0, 0)
	return res
}

// 寻找第i行未出现的数字集合A
func notInRow(nums [9][9]uint, i int) []uint {
	f := map[uint]bool{}
	res := []uint{}
	for _, x := range nums[i] {
		f[x] = true
	}

	var j uint = 1
	for ; j < 10; j++ {
		if _, ok := f[j]; !ok {
			res = append(res, j)
		}
	}
	return res
}

// 寻找第j列未出现的数字集合B
func notInColumn(nums [9][9]uint, j int) []uint {
	f := map[uint]bool{}
	var res []uint
	for _, x := range nums {
		f[x[j]] = true
	}

	var i uint = 1
	for ; i < 10; i++ {
		if _, ok := f[i]; !ok {
			res = append(res, i)
		}
	}
	return res
}

// 寻找第i行第j列所在子方块未出现的数字集合C
func notInSub(nums [9][9]uint, i, j int) []uint {
	f := map[uint]bool{}
	var res []uint
	for rlow, rhigh := i/3*3, i/3*3+3; rlow < rhigh; rlow++ {
		for clow, chigh := j/3*3, j/3*3+3; clow < chigh; clow++ {
			f[nums[rlow][clow]] = true
		}
	}

	var x uint = 1
	for ; x < 10; x++ {
		if _, ok := f[x]; !ok {
			res = append(res, x)
		}
	}
	return res
}

// 求集合的交集
func intersection(a ...[]uint) []uint {
	f := map[uint]uint{}
	res := []uint{}
	for _, set := range a {
		for _, n := range set {
			f[n]++
		}
	}

	for k, v := range f {
		if v == 3 {
			res = append(res, k)
		}
	}

	return res
}

// 二维数组中0的数量
func zeroCount(nums [9][9]uint) uint {
	var res int
	for _, row := range nums {
		for _, n := range row {
			if n == 0 {
				res++
			}
		}
	}
	return uint(res)
}

func printMatrix(nums [9][9]uint) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(nums[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	tests := []struct {
		nums [9][9]uint
		res  [9][9]uint
	}{
		// {
		// 	[9][9]uint{
		// 		{0, 9, 2, 4, 8, 1, 7, 6, 3},
		// 		{4, 1, 3, 7, 6, 2, 9, 8, 5},
		// 		{8, 6, 7, 3, 5, 9, 4, 1, 2},
		// 		{6, 2, 4, 1, 9, 5, 3, 7, 8},
		// 		{7, 5, 9, 8, 4, 3, 1, 2, 6},
		// 		{1, 3, 8, 6, 2, 7, 5, 9, 4},
		// 		{2, 7, 1, 5, 3, 8, 6, 4, 9},
		// 		{3, 8, 6, 9, 1, 4, 2, 5, 7},
		// 		{0, 4, 5, 2, 7, 6, 8, 3, 1},
		// 	},
		// 	[9][9]uint{
		// 		{5, 9, 2, 4, 8, 1, 7, 6, 3},
		// 		{4, 1, 3, 7, 6, 2, 9, 8, 5},
		// 		{8, 6, 7, 3, 5, 9, 4, 1, 2},
		// 		{6, 2, 4, 1, 9, 5, 3, 7, 8},
		// 		{7, 5, 9, 8, 4, 3, 1, 2, 6},
		// 		{1, 3, 8, 6, 2, 7, 5, 9, 4},
		// 		{2, 7, 1, 5, 3, 8, 6, 4, 9},
		// 		{3, 8, 6, 9, 1, 4, 2, 5, 7},
		// 		{9, 4, 5, 2, 7, 6, 8, 3, 1},
		// 	},
		// },
		{
			[9][9]uint{
				{0, 0, 8, 7, 1, 9, 2, 4, 5},
				{9, 0, 5, 2, 3, 4, 0, 8, 6},
				{0, 7, 4, 8, 0, 6, 1, 0, 3},
				{7, 0, 3, 0, 9, 2, 0, 0, 0},
				{5, 0, 0, 0, 0, 0, 0, 0, 0},
				{8, 6, 1, 4, 0, 3, 5, 2, 9},
				{4, 0, 0, 0, 2, 0, 0, 0, 8},
				{0, 0, 0, 0, 0, 0, 0, 7, 0},
				{1, 0, 7, 0, 6, 8, 0, 5, 0},
			},
			[9][9]uint{
				{6, 3, 8, 7, 1, 9, 2, 4, 5},
				{9, 1, 5, 2, 3, 4, 7, 8, 6},
				{2, 7, 4, 8, 5, 6, 1, 9, 3},
				{7, 4, 3, 5, 9, 2, 8, 6, 1},
				{5, 9, 2, 6, 8, 1, 4, 3, 7},
				{8, 6, 1, 4, 7, 3, 5, 2, 9},
				{4, 5, 6, 3, 2, 7, 9, 1, 8},
				{3, 8, 9, 1, 4, 5, 6, 7, 2},
				{1, 2, 7, 9, 6, 8, 3, 5, 4},
			},
		},
		// {
		// 	[9][9]uint{
		// 		{0, 9, 5, 0, 2, 0, 0, 6, 0},
		// 		{0, 0, 7, 1, 0, 3, 9, 0, 2},
		// 		{6, 0, 0, 0, 0, 5, 3, 0, 4},
		// 		{0, 4, 0, 0, 1, 0, 6, 0, 7},
		// 		{5, 0, 0, 2, 0, 7, 0, 0, 9},
		// 		{7, 0, 3, 0, 9, 0, 0, 2, 0},
		// 		{0, 0, 9, 8, 0, 0, 0, 0, 6},
		// 		{8, 0, 6, 3, 0, 2, 1, 0, 5},
		// 		{0, 5, 0, 0, 7, 0, 2, 8, 3},
		// 	},
		// 	[9][9]uint{
		// 		{3, 9, 5, 7, 2, 4, 8, 6, 1},
		// 		{4, 8, 7, 1, 6, 3, 9, 5, 2},
		// 		{6, 2, 1, 9, 8, 5, 3, 7, 4},
		// 		{9, 4, 2, 5, 1, 8, 6, 3, 7},
		// 		{5, 6, 8, 2, 3, 7, 4, 1, 9},
		// 		{7, 1, 3, 4, 9, 6, 5, 2, 8},
		// 		{2, 3, 9, 8, 5, 1, 7, 4, 6},
		// 		{8, 7, 6, 3, 4, 2, 1, 9, 5},
		// 		{1, 5, 4, 6, 7, 9, 2, 8, 3},
		// 	},
		// },
	}

	for _, test := range tests {
		fmt.Println("input:")
		printMatrix(test.nums)
		fmt.Println("---------------------------------")
		res := Sodoku(test.nums)
		fmt.Println("output:")
		printMatrix(res)

		if !reflect.DeepEqual(res, test.res) {
			fmt.Println("failed,want:")
			printMatrix(test.res)
		} else {
			fmt.Println("passed")
		}
	}

}
