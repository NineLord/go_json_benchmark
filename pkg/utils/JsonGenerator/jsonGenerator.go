package JsonGenerator

import (
	"fmt"
	"github.com/NineLord/go_json_benchmark/pkg/utils/JsonType"
	"github.com/NineLord/go_json_benchmark/pkg/utils/Randomizer"
	"reflect"
	"strings"
)

type jsonGenerator struct {
	characterPoll    []rune
	numberOfLetters  uint
	depth            uint
	numberOfChildren uint
}

func newJsonGenerator(characterPoll string, numberOfLetters uint, depth uint, numberOfChildren uint) *jsonGenerator {
	result := makeJsonGenerator(characterPoll, numberOfLetters, depth, numberOfChildren)
	return &result
}

func makeJsonGenerator(characterPoll string, numberOfLetters uint, depth uint, numberOfChildren uint) jsonGenerator {
	return jsonGenerator{
		characterPoll:    []rune(characterPoll),
		numberOfLetters:  numberOfLetters,
		depth:            depth,
		numberOfChildren: numberOfChildren,
	}
}

func GenerateJson(characterPoll string, numberOfLetters uint, depth uint, numberOfChildren uint) (map[string]interface{}, error) {
	return newJsonGenerator(characterPoll, numberOfLetters, depth, numberOfChildren).generateFullTree()
}

func (jsonGenerator *jsonGenerator) generateFullTree() (map[string]interface{}, error) {
	result := make(map[string]interface{})

	currentNodes := make([]interface{}, 1)
	currentNodes = append(currentNodes, result)
	nextLevelNodes := make([]interface{}, 0)
	lastLevel := jsonGenerator.depth - 1

	for level := uint(0); level < lastLevel; level++ {

		for len(currentNodes) != 0 {
			var currentNode interface{}
			currentNode, currentNodes = currentNodes[len(currentNodes)-1], currentNodes[:len(currentNodes)-1] // Pop
			currentNodeType := reflect.TypeOf(currentNode).Kind()

			switch currentNodeType {
			case reflect.Slice:
				currentNode := currentNode.([]interface{})
				for _nodeCount := uint(0); _nodeCount < jsonGenerator.numberOfChildren; _nodeCount++ {
					var childNodeValue interface{}
					if level == lastLevel {
						childNodeValue = JsonType.GetRandomLeafJson()
					} else {
						childNodeValue = JsonType.GetRandomNoneLeafJson(jsonGenerator.numberOfChildren)
					}
					currentNode = append(currentNode, childNodeValue)
					nextLevelNodes = append(nextLevelNodes, childNodeValue)
				}
			case reflect.Map:
				currentNode := currentNode.(map[string]interface{})
				for _nodeCount := uint(0); _nodeCount < jsonGenerator.numberOfChildren; _nodeCount++ {
					var childNodeValue interface{}
					if level == lastLevel {
						childNodeValue = JsonType.GetRandomLeafJson()
					} else {
						childNodeValue = JsonType.GetRandomNoneLeafJson(jsonGenerator.numberOfChildren)
					}
					currentNode[jsonGenerator.getRandomNodeName()] = childNodeValue
					nextLevelNodes = append(nextLevelNodes, childNodeValue)
				}
			default:
				panic(fmt.Sprintf("generateFullTree unknown JSON type: %d", currentNodeType))
			}
		}

		currentNodes = nextLevelNodes
		nextLevelNodes = make([]interface{}, 0)
	}

	return result, nil
}

// #region Helper methods
func (jsonGenerator *jsonGenerator) getRandomNodeCharacter() *rune {
	index := Randomizer.GetRandomIndexFromArray(jsonGenerator.characterPoll)
	char := jsonGenerator.characterPoll[index]
	return &char
}

func (jsonGenerator *jsonGenerator) getRandomNodeName() string {
	var stringBuilder strings.Builder
	for count := uint(0); count < jsonGenerator.numberOfLetters; count++ {
		stringBuilder.WriteRune(*jsonGenerator.getRandomNodeCharacter())
	}
	return stringBuilder.String()
}

// #endregion
