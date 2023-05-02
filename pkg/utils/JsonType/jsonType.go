package JsonType

import (
	"fmt"
	"github.com/NineLord/go_json_benchmark/pkg/utils/Randomizer"
	"strings"
)

type ValueNonLeafType uint

const (
	Array ValueNonLeafType = iota
	Object
)

type ValueLeafType uint

const (
	Null ValueLeafType = iota
	Bool
	Number
	String
)

var VariantsValueLeafTypes = [4]ValueLeafType{Null, Bool, Number, String}

// Generic Type: [ValueTypes valueTypes]
type valueTypes interface {
	ValueNonLeafType | ValueLeafType
}

var ALPHABET = []rune("AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz")

func getRandomNoneLeafJsonType() ValueNonLeafType {
	if Randomizer.GetRandomBoolean() {
		return Array
	} else {
		return Object
	}
}

func GetRandomNoneLeafJson(numberOfChildren uint) interface{} {
	jsonType := getRandomNoneLeafJsonType()
	switch jsonType {
	case Array:
		return make([]interface{}, numberOfChildren, numberOfChildren)
	case Object:
		return make(map[string]interface{})
	default:
		panic(fmt.Sprintf("GetRandomNoneLeafJson unknown JSON type: %d", jsonType))
	}
}

func getRandomLeafJsonType() ValueLeafType {
	return *Randomizer.GetRandomValueFromArray(VariantsValueLeafTypes[:])
}

func GetRandomLeafJson() interface{} {
	jsonType := getRandomLeafJsonType()
	switch jsonType {
	case Null:
		return nil
	case Bool:
		return Randomizer.GetRandomBoolean()
	case Number:
		return Randomizer.GetRandomNumberInRangeFloat64(-1_000_000_000.0, 1_000_000_000.0)
	case String:
		numberOfLetters := Randomizer.GetRandomNumberInRangeInt(0, 32+1)
		var stringBuilder strings.Builder
		for _count := 0; _count < numberOfLetters; _count++ {
			stringBuilder.WriteRune(*Randomizer.GetRandomValueFromArray(ALPHABET))
		}
		return stringBuilder.String()
	default:
		panic(fmt.Sprintf("GetRandomLeafJson unknown JSON type: %d", jsonType))
	}
}
