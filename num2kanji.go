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
	TEN                  = 0
	HUNDRED              = 1
	THOUSAND             = 2
	TEN_THOUSAND         = 3
	ONE_HUNDRED_THOUSAND = 4
	ONE_TRILLION         = 5
)

const LIMIT_LOOP int = 6

var unitsByKanji = map[int]string{
	TEN:                  "十",
	HUNDRED:              "百",
	THOUSAND:             "千",
	TEN_THOUSAND:         "万",
	ONE_HUNDRED_THOUSAND: "億",
	ONE_TRILLION:         "兆",
}

// getUnit this func's param is unit int, return string
func getUnit(unit int) string {
	return unitsByKanji[unit]
}

func Kanji2number(param string) (result string, err error) {
	num, err := strconv.Atoi(param)
	if err != nil {
		result = ""
		return
	}

	result = recursion2kanji(num, "")
	return
}

func recursion2kanji(num int, result string) string {
	for i := LIMIT_LOOP; i >= 0; i-- {
		bias := math.Pow(10, float64(i))
		if num >= int(bias) {
			top := math.Floor(float64(num / int(bias)))
			if int(top) >= 10 {
				recursion2kanji(int(top), result)
			} else {
				result += toKanji(int(top))
			}
			result += getUnit(i)
			num -= int(top) * int(bias)
		}
	}
	result += toKanji(num)
	return result
}
