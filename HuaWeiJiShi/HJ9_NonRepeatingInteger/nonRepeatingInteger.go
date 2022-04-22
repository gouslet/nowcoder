/*
 * File: \HuaWeiJiShi\HJ9 提取不重复的整数\nonRepeatingInteger.go                      *
 * Project: newcoder                                                           *
 * Created At: Wednesday, 2022/04/20 , 10:55:19                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/04/20 , 11:39:29                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

func NonRepeatingInteger1(n int) int {
	digit := []int{}
	exists := map[int]bool{}
	for n != 0 {
		if x := n % 10; exists[x] == false {
			digit = append(digit, x)
			exists[x] = true
		}
		n /= 10
	}
	for i, l := 0, len(digit); i < l; i++ {
		n *= 10
		n += digit[i]
	}
	return n
}

func NonRepeatingInteger2(n int) int {
	exists := map[int]bool{}
	res := 0
	for n != 0 {
		if x := n % 10; exists[x] == false {
			res = 10*res + x
			exists[x] = true
		}
		n /= 10
	}
	return res
}

func NonRepeatingInteger3(n int) int {
	exists := [10]uint{}
	res := 0
	for n != 0 {
		if x := n % 10; exists[x] == 0 {
			res = 10*res + x
			exists[x] = 1
		}
		n /= 10
	}
	return res
}

func main() {
	fmt.Println(NonRepeatingInteger2(1122334))
}
