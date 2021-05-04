package main

import (
	"errors"
	"math"
	"strings"
)

var kanji2number = map[string]int{
	KANJI_ZERO:  ZERO,
	KANJI_ONE:   ONE,
	KANJI_TWO:   TWO,
	KANJI_THREE: THREE,
	KANJI_FOUR:  FOUR,
	KANJI_FIVE:  FIVE,
	KANJI_SIX:   SIX,
	KANJI_SEVEN: SEVEN,
	KANJI_EIGHT: EIGHT,
	KANJI_NINE:  NINE,
}

// 降順であること
var unitsByKanjiUnder = map[string]int{
	KANJI_THOUSAND: THOUSAND,
	KANJI_HUNDRED:  HUNDRED,
	KANJI_TEN:      TEN,
}

// 降順であること
var unitsByKanjiOver = map[string]int{
	KANJI_ONE_TRILLION:         ONE_TRILLION,
	KANJI_ONE_HUNDRED_THOUSAND: ONE_HUNDRED_THOUSAND,
	KANJI_TEN_THOUSAND:         TEN_THOUSAND,
}

func toInt(str string) (result int, canEx bool) {
	result, canEx = kanji2number[str]
	return
}

func kanji2num(param string) (result int, err error) {
	if param == "" {
		err = errors.New("invalid param")
		return 0, err
	}

	result, canEx := recursion2numOver(param)
	if !canEx {
		err = errors.New("invalid param")
		return 0, err
	}
	return
}

// 万以下
func recursion2numUnder(numstr string) (result int, canEx bool) {
	// 10~
	for key, value := range unitsByKanjiUnder {
		strSlice := strings.Split(numstr, key)
		top, canEx := toInt(strSlice[0])
		if !canEx {
			return 0, canEx
		} else {
			result += top * int(math.Pow10(value))
			numstr = strSlice[1]
		}
	}

	// ~9
	num, canEx := toInt(numstr)
	if !canEx {
		return 0, canEx
	} else {
		result += num
	}
	return
}

// 万以上
func recursion2numOver(numstr string) (result int, canEx bool) {
	// 万以上
	for key, value := range unitsByKanjiOver {
		strSlice := strings.Split(numstr, key)
		top, canEx := recursion2numUnder(strSlice[0])
		if !canEx {
			return 0, canEx
		} else {
			result += top * int(math.Pow10(value))
			numstr = strSlice[1]
		}
	}

	// 万以下
	top, canEx := recursion2numUnder(numstr)
	if !canEx {
		return 0, canEx
	} else {
		result += top
	}
	return
}
