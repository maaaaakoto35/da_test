package main

import (
	"math"
	"strconv"
)

// constant 0~9 number
const (
	ZERO  = 0
	ONE   = 1
	TWO   = 2
	THREE = 3
	FOUR  = 4
	FIVE  = 5
	SIX   = 6
	SEVEN = 7
	EIGHT = 8
	NINE  = 9
)

var number2kanji = map[int]string{
	ZERO:  "零",
	ONE:   "壱",
	TWO:   "弐",
	THREE: "参",
	FOUR:  "四",
	FIVE:  "五",
	SIX:   "六",
	SEVEN: "七",
	EIGHT: "八",
	NINE:  "九",
}

// toKanji this func's param is num int, return string
func toKanji(num int) string {
	return number2kanji[num]
}

// constant units
const (
	TEN                  = 1
	HUNDRED              = 2
	THOUSAND             = 3
	TEN_THOUSAND         = 4
	ONE_HUNDRED_THOUSAND = 8
	ONE_TRILLION         = 12
)

const LIMIT_LOOP int = 5
const MAX = 9999999999999999

var unitsByKanji = map[int]string{
	TEN:                  "拾",
	HUNDRED:              "百",
	THOUSAND:             "千",
	TEN_THOUSAND:         "万",
	ONE_HUNDRED_THOUSAND: "億",
	ONE_TRILLION:         "兆",
}

var unitsByInt [LIMIT_LOOP + 1]int = [LIMIT_LOOP + 1]int{1, 2, 3, 4, 8, 12}

// getUnit this func's param is unit int, return string
func getUnitByKanji(unit int) string {
	return unitsByKanji[unit]
}

func Kanji2number(param string) (result string, err error) {
	num, err := strconv.Atoi(param)
	if err != nil {
		result = ""
		return
	}
	if num == ZERO {
		result = toKanji(num)
	} else if num < ZERO || num > MAX {
		result = ""
	} else {
		result = recursion2kanji(num, "")
	}
	return
}

func recursion2kanji(num int, result string) string {
	for i := LIMIT_LOOP; i >= 0; i-- {
		loopUnit := unitsByInt[i]
		bias := math.Pow10(loopUnit)
		if num >= int(bias) {
			top := math.Floor(float64(num / int(bias)))
			if int(top) >= 10 {
				result += recursion2kanji(int(top), result)
			} else {
				result += toKanji(int(top))
			}
			result += getUnitByKanji(loopUnit)
			num -= int(top) * int(bias)
		}
	}
	if num != ZERO {
		result += toKanji(num)
	}
	return result
}
