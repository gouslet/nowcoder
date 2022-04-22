/*
 * File: \HuaWeiJiShi\HJ94_VoteCounting\voteCounting.go                        *
 * Project: newcoder                                                           *
 * Created At: Wednesday, 2022/04/20 , 12:05:14                                *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Wednesday, 2022/04/20 , 13:35:37                             *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import "fmt"

func VoteCounting(votes map[string]uint, ss []string) uint {
	var valid uint
	var voteCounting func(string) bool

	voteCounting = func(str string) bool {
		l := len(str)
		if l == 0 {
			return true
		}
		for i := 1; i <= l; i++ {
			subs := str[:i]
			if _, ok := votes[subs]; ok {
				if voteCounting(str[i:]) == true {
					votes[subs]++
					return true
				}
			}
		}

		return false
	}

	for _, s := range ss {
		if voteCounting(s) == false {
			valid++
		}
	}
	return valid
}

func VoteCounting2(votes map[string]uint, ss []string) uint {
	var valid uint
	var voteCounting func(string, int, int) bool
	voteCounting = func(str string, i, l int) bool {
		if i == l {
			return true
		}
		for j := 1; j <= l; j++ {
			subs := str[i:j]
			if _, ok := votes[subs]; ok {
				if voteCounting(str, j, l) == true {
					votes[subs]++
					return true
				}
			}
		}

		return false
	}

	for _, s := range ss {
		l := len(s)
		if voteCounting(s, 0, l) == false {
			valid++
		}
	}
	return valid
}

func main() {
	votes := map[string]uint{
		"A": 0,
		"B": 0,
		"C": 0,
		"D": 0,
	}
	ss := []string{
		"A",
		"B",
		"AC",
		"BD",
		"DE",
		"FE",
		"AGB",
	}

	fmt.Println(VoteCounting2(votes, ss))
	fmt.Println(votes)
}
