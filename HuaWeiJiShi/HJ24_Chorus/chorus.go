/*
 * File: \HuaWeiJiShi\HJ24_Chorus\chorus.go                                    *
 * Project: newcoder                                                           *
 * Created At: Thursday, 2022/04/21 , 18:14:38                                 *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Saturday, 2022/04/23 , 00:31:12                              *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"fmt"
	"math"
	"runtime"
)

func Chorus1(heights []uint) uint {
	l := len(heights)

	dp := make([]uint, l)
	dp[0] = 0
	var i, t int
	var temp uint
	flag := make([]uint, l)
	flag[0] = 0

	for i = 1; i < l; i++ {
		dp[i] = math.MaxUint
		if heights[i] > heights[0] {
			flag[i] = 1
		}
		for t = 1; t < i; t++ {
			if heights[t] > heights[i] {
				if temp = uint(i) - flag[t] - 1; flag[t] > 0 && temp < dp[i] {
					dp[i] = temp
				}
				if temp = uint(i) - uint(t) + dp[t] - 1; uint(t)-dp[t] > 1 && temp < dp[i] {
					dp[i] = temp
				}
			} else if temp = flag[t] + 1; heights[t] < heights[i] && flag[i] < temp {
				flag[i] = temp
			}
		}
		// fmt.Printf("%d: %d\n", i, flag[i])
		if dp[i] == math.MaxUint || dp[i-1]+1 < dp[i] {
			dp[i] = dp[i-1] + 1
		}
		fmt.Printf("%d: %d\n", i, dp[i])
	}

	return dp[l-1]
}

// Time:O(n^3) Space:O(n)
func Chorus2(heights []uint) uint {
	l := len(heights) - 1

	var res uint = math.MaxUint
	d := func(k int) uint { //以heights[k]为中心的合唱队需出列的最小值
		lis := make([]uint, l+1)
		lis[0] = 0
		lis[l] = 0
		for i := 1; i <= k; i++ {
			for j := 0; j < i; j++ {
				if heights[i] > heights[j] && lis[i] < lis[j]+1 {
					lis[i] = lis[j] + 1
				}
			}
		}
		// fmt.Printf("%d: left %v\n", k, lis)
		left := lis[k] //递增序列heights[i:k)的最大长度

		lis[k] = 0
		for i := l - 1; i >= k; i-- {
			for j := i + 1; j < l; j++ {
				if heights[i] > heights[j] && lis[i] < lis[j]+1 {
					lis[i] = lis[j] + 1
				}
			}
		}
		// fmt.Printf("%d: right %v\n", k, lis)

		right := lis[k] //递减序列heights(k:l]的最大长度
		// fmt.Printf("k: %d l :%d left: %d right: %d\n", k, l, left, right)
		if left == 0 || right == 0 {
			return uint(l)
		}
		return uint(l) - left - right
	}
	if l == 0 {
		return 0
	}
	for i := 1; i < l; i++ {
		if d(i) < res {
			res = d(i)
		}
	}
	return res
}

// Time:O(n^2) Space:O(n)
func Chorus3(heights []uint) uint {
	l := len(heights) - 1
	if l == 0 {
		return 0
	}

	var res, temp, left, right uint = math.MaxUint, 0, 0, 0
	var j, t int

	lis := make([]uint, l+1) //lis[i]--以i位置元素为末尾元素的最长递增子序列长度
	lds := make([]uint, l+1) //lds[i]--以i位置元素为首元素的最长递减子序列长度
	lis[0], lds[0] = 1, 1

	for i := 1; i < l+1; i++ {
		lis[i] = 1
		for j := 0; j < i; j++ { //lis[i] = max{lis[j]+1 | heights[i] > heights[j],lis[j]+1 > lis[i]}
			if tmp := lis[j] + 1; heights[i] > heights[j] && lis[i] < tmp {
				lis[i] = tmp
			}
		}
		t = l - i
		lds[t] = 1
		for j = l; j > t; j-- { //lis[t] = max{lis[j]+1 | heights[t] > heights[j],lis[j]+1 > lis[t]}
			if tmp := lds[j] + 1; heights[t] > heights[j] && lds[t] < tmp {
				lds[t] = tmp
			}
		}
	}
	for k := 1; k < l; k++ { //分别以1~l-1位置的每个人为中心，找出合唱队最大长度，即需出列的同学数量最少
		fmt.Printf("%d: left %v\n", k, lis)
		left = lis[k] //递增序列heights[i:k]的最大长度

		// fmt.Printf("%d: right %v\n", k, lis)

		right = lds[k] //递减序列heights[k:l]的最大长度
		// fmt.Printf("k: %d l :%d left: %d right: %d\n", k, l, left, right)

		if left == 1 || right == 1 { //以k位置元素为末尾元素的递增序列长度为1，或者以k位置元素为首元素的递减序列长度为1，则无法形成合唱队
			temp = uint(l) // 此时合唱队仅仅一人
		} else {
			temp = uint(l) - left - right + 2 //l+1 - (左侧最长递增子序列与右侧最长递减子序列之和-1)，中间一人被计算了两次，需减去1
		}

		if temp < res {
			res = temp
		}

	}
	return res
}

// Time:O(n^logn) Space:O(n)
func Chorus4(heights []uint) uint {
	l := len(heights) - 1

	var res, temp, left, right uint = math.MaxUint, 0, 0, 0
	var t int

	lis := make([]uint, l+1) //左侧最长递增子序列长度
	lds := make([]uint, l+1) //右侧最长递减子序列长度

	if l == 0 {
		return 0
	}
	ltails := make([]uint, l+1) //ltails[k]，当前长度为k的左侧最长递增子序列的末尾元素
	rtails := make([]uint, l+1) //rtails[k]，当前长度为k的右侧最长递减子序列的首元素
	leftNumbers := 0
	rightNumbers := 0

	for i := 0; i < l+1; i++ {
		start, end := 0, leftNumbers
		for start < end {
			mid := start + (end-start)>>1
			if ltails[mid] < heights[i] {
				start = mid + 1
			} else {
				end = mid
			}
		}
		ltails[start] = heights[i]
		if start == leftNumbers {
			leftNumbers++
		}
		lis[i] = uint(leftNumbers)

		t = l - i
		start, end = 0, rightNumbers
		for start < end {
			mid := start + (end-start)>>1
			if rtails[mid] < heights[t] {
				start = mid + 1
			} else {
				end = mid
			}
		}
		rtails[start] = heights[t]
		if start == rightNumbers {
			rightNumbers++
		}
		lds[t] = uint(rightNumbers)
	}
	// fmt.Println("lis: ", lis)
	// fmt.Println("lds: ", lds)
	for k := 1; k < l; k++ {
		// fmt.Printf("%d: left %v\n", k, lis)
		left = lis[k] //递增序列heights[i:k]的最大长度

		// fmt.Printf("%d: right %v\n", k, lis)

		right = lds[k] //递减序列heights[k:l]的最大长度
		// fmt.Printf("k: %d l :%d left: %d right: %d\n", k, l, left, right)

		if left == 1 || right == 1 {
			temp = uint(l)
		} else {
			temp = uint(l) - left - right + 2
		}

		if temp < res {
			res = temp
		}

	}
	return res
}

// func left(nums []uint) {
// 	l := len(nums)
// 	lis := make([]uint, l)
// 	lis[0] = 0
// 	for i := 1; i < l; i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] && lis[i] < lis[j]+1 {
// 				lis[i] = lis[j] + 1
// 			}
// 		}
// 	}
// 	fmt.Println(lis)
// }

// func right(nums []uint, k int) {
// 	l := len(nums)
// 	lis := make([]uint, l-k)
// 	lis[l-k-1] = 0
// 	for i := l - k - 2; i > -1; i-- {
// 		for j := i + 1; j < l-k; j++ {
// 			if nums[i] > nums[j] && lis[i] < lis[j]+1 {
// 				lis[i] = lis[j] + 1
// 			}
// 		}
// 	}
// 	fmt.Println(lis)
// }

func main() {
	heights := [][]uint{
		// {
		// 	186, 186, 150, 200, 160, 130, 197, 200,
		// },
		{
			// 186, 186, 150, 200, 160, 130, 197, 200,
			186, 186, 186, 156,
		},
		// {
		// 	16, 103, 132, 23, 211, 75, 155, 82, 32, 48, 79, 183, 13, 91, 51, 172, 109, 102, 189, 121, 12, 120, 116, 133, 79, 120, 116, 208, 47, 110, 65, 187, 69, 143, 140, 173, 203, 35, 184, 49, 245, 50, 179, 63, 204, 34, 218, 11, 205, 100, 90, 19, 145, 203, 203, 215, 72, 108, 58, 198, 95, 116, 125, 235, 156, 133, 220, 236, 125, 29, 235, 170, 130, 165, 155, 54, 127, 128, 204, 62, 59, 226, 233, 245, 46, 3, 14, 108, 37, 94, 52, 97, 159, 190, 143, 67, 24, 204, 39, 222, 245, 233, 11, 80, 166, 39, 224, 12, 38, 13, 85, 21, 47, 25, 180, 219, 140, 201, 11, 42, 110, 209, 77, 136,
		// },
		{
			16, 103, 132, 23, 211, 75, 155, 82, 32, 48, 79, 183, 13, 91, 51,
		},
	}
	for _, h := range heights {
		fmt.Printf("Chorus(%v) = %d\n", h, Chorus3(h))
		// left(h)
		// for i := 1; i < len(h); i++ {
		// 	right(h, i)
		// }
	}
	fmt.Println(runtime.Version())
}
