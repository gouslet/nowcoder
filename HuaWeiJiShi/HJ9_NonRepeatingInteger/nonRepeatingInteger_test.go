/*
 * File: \HuaWeiJiShi\HJ44_Sodoku\sodoku_test.go                               *
 * Project: newcoder                                                           *
 * Created At: Tuesday, 2022/04/19 , 23:20:08                                  *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/04/20 , 11:41:37                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */

package main

import (
	"testing"
)

func TestNonRepeatingInteger(t *testing.T) {
	funcs := map[string]func(int) int{
		"NonRepeatingInteger1": NonRepeatingInteger1,
		"NonRepeatingInteger2": NonRepeatingInteger2,
		"NonRepeatingInteger3": NonRepeatingInteger3,
	}
	tests := []struct {
		n, res int
	}{
		{1, 1},
		{10, 1},
		{9765, 5679},
		{1122334, 4321},
		{10e8, 1},
	}
	for fn, f := range funcs {
		for _, test := range tests {
			if res := f(test.n); res != test.res {
				t.Errorf("%s(%d) = %d,want %d\n", fn, test.n, res, test.res)
			}
		}
	}

}
