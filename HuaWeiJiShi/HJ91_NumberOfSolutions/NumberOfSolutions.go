/*
 * File: \HuaWeiJiShi\HJ91_NumberOfSolutions\NumberOfSolutions.go              *
 * Project: newcoder                                                           *
 * Created At: Wednesday, 2022/04/20 , 14:52:41                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/04/20 , 15:02:24                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import "fmt"

func NumberOfSolutions(n, m byte) byte {
	nums := make([]byte, m)
	nums[0] = 1
	var i byte
	var j byte
	for ; i < n; i++ {
		for j = 0; j < m; j++ {
			if j > 0 {
				nums[j] += nums[j-1]
			}
		}
	}
	return nums[m-1]
}

func main() {
	fmt.Println(NumberOfSolutions(3, 3))
}
