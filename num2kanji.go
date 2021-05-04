package main

import (
	"math"
	"strconv"
)

var number2kanji = map[int]string{
	ZERO:  KANJI_ZERO,
	ONE:   KANJI_ONE,
	TWO:   KANJI_TWO,
	THREE: KANJI_THREE,
	FOUR:  KANJI_FOUR,
	FIVE:  KANJI_FIVE,
	SIX:   KANJI_SIX,
	SEVEN: KANJI_SEVEN,
	EIGHT: KANJI_EIGHT,
	NINE:  KANJI_NINE,
}

// toKanji this func's param is num int, return string
func toKanji(num int) string {
	return number2kanji[num]
}

const LIMIT_LOOP int = 5
const MAX = 9999999999999999

var unitsByIntMap = map[int]string{
	TEN:                  KANJI_TEN,
	HUNDRED:              KANJI_HUNDRED,
	THOUSAND:             KANJI_THOUSAND,
	TEN_THOUSAND:         KANJI_TEN_THOUSAND,
	ONE_HUNDRED_THOUSAND: KANJI_ONE_HUNDRED_THOUSAND,
	ONE_TRILLION:         KANJI_ONE_TRILLION,
}

var unitsByInt [LIMIT_LOOP + 1]int = [LIMIT_LOOP + 1]int{TEN, HUNDRED, THOUSAND, TEN_THOUSAND, ONE_HUNDRED_THOUSAND, ONE_TRILLION}

// getUnit this func's param is unit int, return string
func getUnitByKanji(unit int) string {
	return unitsByIntMap[unit]
}

func num2kanji(param string) (result string, err error) {
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
		result = recursion2kanji(num)
	}
	return
}

func recursion2kanji(num int) string {
	var result = ""
	for i := LIMIT_LOOP; i >= 0; i-- {
		loopUnit := unitsByInt[i]
		bias := math.Pow10(loopUnit)
		if num >= int(bias) {
			top := math.Floor(float64(num / int(bias)))
			if int(top) >= 10 {
				result += recursion2kanji(int(top))
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
