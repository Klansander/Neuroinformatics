package main

import (
	"fmt"
	"math/rand"
	"time"
)

func check_randNum(randNums [9]int, randNum int) bool {
	for i := 0; i < 9; i++ {
		if randNum == randNums[i] {
			return false
		}
	}
	return true
}
func init_random(mass *[300][9]int, er int) {
	for i := 0; i < 300; i++ {
		randNums := [9]int{9, 9, 9, 9, 9, 9, 9, 9, 9}
		for k := 0; k < er; k++ {
			randNum := 0
			for {
				rand.Seed(time.Now().UnixNano())
				randNum = 0 + rand.Intn(9-0+1)
				if check_randNum(randNums, randNum) {
					break
				}
			}
			randNums[k] = randNum
			if mass[i][randNum] == 0 {
				mass[i][randNum] = 1
			} else {
				mass[i][randNum] = 0
			}

		}
	}

}
func massIf(mass1 [9]int, mass2 [9]int) bool {
	for i := 0; i < 9; i++ {
		if mass1[i] != mass2[i] {
			return false
		}
	}
	return true
}
func counter(mass [9]int, w *[9]int, sum bool) bool {
	if sum {
		for i := 0; i < 9; i++ {
			if mass[i] == 1 {
				w[i]++
			}
		}
	} else {
		for i := 0; i < 9; i++ {
			if mass[i] == 1 {
				w[i]--
			}
		}
	}

	return false
}
func comparison(mass [9]int, mass2 [30][9]int) bool {
	var k, m int
	for j := 0; j < 30; j++ {
		k = 0
		for i := 0; i < 9; i++ {
			if mass[i] != mass2[j][i] {
				k++
			}
		}
		if k == 0 {
			m++
		}
	}
	if m != 0 {
		return true
	} else {
		return false
	}
}
func Mul(mass [9]int, w [9]int) int {
	sum := 0
	for i := 0; i < 9; i++ {
		if mass[i] == 1 {
			sum = sum + w[i]
		}
	}
	return sum
}
func result(mass [300][9]int, w *[9]int, Q int, start [30][9]int) (float64, int) {
	learn := false
	var iter2, izm int
	res := 0.0
	count := 0
	mass_ob := [300]int{}
	for !learn {
		learn = true
		izm = 0
		for k := 0; k < 300; k++ {
			num := Mul(mass[k], *w)
			if comparison(mass[k], start) { 
				if num <= Q {
					learn = counter(mass[k], w, true)
					izm++
				}
			} else {
				if num > Q {
					learn = counter(mass[k], w, false)
					izm++
				}
			}
		}
		mass_ob[iter2] = izm
		if iter2 >= 1 { 
			if mass_ob[iter2-0]-mass_ob[iter2-1] == 1 || mass_ob[iter2-0]-mass_ob[iter2-1] == -1 || mass_ob[iter2-0]-mass_ob[iter2-1] == 2 || mass_ob[iter2-0]-mass_ob[iter2-1] == -2 { // or learn
				count = mass_ob[iter2]
				res = float64(count) / 300.0
				break
			} else if mass_ob[iter2] == 0 {
				res = 0.0
				break
			}

		}
		iter2++

	}
	return res, count
}

func main() {
	mass := [300][9]int{
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
		{1, 1, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 0, 0, 1},
		{0, 1, 1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 1, 1},
		{0, 0, 1, 0, 1, 1, 0, 1, 1},
		{0, 1, 0, 1, 0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 0, 1, 1, 0, 1, 0, 0},
	}
	mass_comparison0 := [30][9]int{}
	mass_comparison1 := [30][9]int{}
	mass_comparison2 := [30][9]int{}
	mass_comparison3 := [30][9]int{}
	mass_comparison4 := [30][9]int{}
	mass_comparison5 := [30][9]int{}
	mass_comparison6 := [30][9]int{}
	mass_comparison7 := [30][9]int{}
	mass_comparison8 := [30][9]int{}
	mass_comparison9 := [30][9]int{}
	startNums := [30]int{}
	priznaki_null := [9]int{1, 1, 0, 1, 0, 1, 0, 1, 1}
	mass1 := [300][9]int{}
	mass2 := [300][9]int{}
	mass3 := [300][9]int{}
	mass4 := [300][9]int{}
	mass5 := [300][9]int{}
	mass6 := [300][9]int{}
	mass7 := [300][9]int{}
	mass8 := [300][9]int{}
	mass9 := [300][9]int{}
	for i := 0; i < 300; i++ {
		for j := 0; j < 9; j++ {
			mass1[i][j], mass2[i][j], mass3[i][j], mass4[i][j], mass5[i][j], mass6[i][j], mass7[i][j], mass8[i][j], mass9[i][j] = mass[i][j], mass[i][j], mass[i][j], mass[i][j], mass[i][j], mass[i][j], mass[i][j], mass[i][j], mass[i][j]
		}
	}
	init_random(&mass1, 1)
	init_random(&mass2, 2)
	init_random(&mass3, 3)
	init_random(&mass4, 4)
	init_random(&mass5, 5)
	init_random(&mass6, 6)
	init_random(&mass7, 7)
	init_random(&mass8, 8)
	init_random(&mass9, 9)
	j := 0
	for i := 0; i < 300; i++ {
		if massIf(mass[i], priznaki_null) {
			startNums[j] = i
			j++
		}
	}
	n := 0
	for i := 0; i < 300; i++ {
		for j := 0; j < 30; j++ {
			if i == startNums[j] {
				for m := 0; m < 9; m++ {
					mass_comparison0[n][m], mass_comparison1[n][m], mass_comparison2[n][m], mass_comparison3[n][m], mass_comparison4[n][m], mass_comparison5[n][m], mass_comparison6[n][m], mass_comparison7[n][m], mass_comparison8[n][m], mass_comparison9[n][m] = mass[i][m], mass1[i][m], mass2[i][m], mass3[i][m], mass4[i][m], mass5[i][m], mass6[i][m], mass7[i][m], mass8[i][m], mass9[i][m]
				}
				n++

			}
		}
	}
	winterference := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference1 := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference2:= [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference3:= [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference4 := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference5 := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference6 := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference7 := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference8 := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	winterference9 := [9]int{3, 2, 7, 9, 5, 6, 2, 7, 3}
	Q := 28
	interference0, count := result(mass, &winterference, Q, mass_comparison0)
	fmt.Printf("???????????????? ?????? 0 ??????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference0, count, winterference)
	interference1, count := result(mass1, &winterference1, Q, mass_comparison1)
	fmt.Printf("???????????????? ?????? 1 ????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference1, count, winterference1)
	interference2, count := result(mass2, &winterference2, Q, mass_comparison2)
	fmt.Printf("???????????????? ?????? 2 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference2, count, winterference2)
	interference3, count := result(mass3, &winterference3, Q, mass_comparison3)
	fmt.Printf("???????????????? ?????? 3 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference3, count, winterference3)
	interference4, count := result(mass4, &winterference4, Q, mass_comparison4)
	fmt.Printf("???????????????? ?????? 4 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference4, count, winterference4)
	interference5, count := result(mass5, &winterference5, Q, mass_comparison5)
	fmt.Printf("???????????????? ?????? 5 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference5, count, winterference5)
	interference6, count := result(mass6, &winterference6, Q, mass_comparison6)
	fmt.Printf("???????????????? ?????? 6 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference6, count, winterference6)
	interference7, count := result(mass7, &winterference7, Q, mass_comparison7)
	fmt.Printf("???????????????? ?????? 7 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference7, count, winterference7)
	interference8, count := result(mass8, &winterference8, Q, mass_comparison8)
	fmt.Printf("???????????????? ?????? 8 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference8, count, winterference8)
	interference9, count := result(mass9, &winterference9, Q, mass_comparison9)
	fmt.Printf("???????????????? ?????? 9 ??????????????:\n\n?????????????????????? ????????????:%0.3f\n???????????????????? ????????????: %d\n???????????????? ?????????????? ?????????? %d\n\n", interference9, count, winterference9)
	start2 := [9]int{0, 0, 1, 0, 1, 0, 1, 0, 0}
	s := Mul(start2, winterference9)
	if s > Q {
		fmt.Println("???????????????? ?????????? ???????????????????????? ?????? 9 ??????????????! ")
	} else {
		fmt.Println("???????????????? ?????????? ???? ???????????????????????? ?????? 9 ??????????????! ")
	}
}

