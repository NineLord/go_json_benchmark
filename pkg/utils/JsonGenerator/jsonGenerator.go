package JsonGenerator

import (
	"fmt"
	"github.com/NineLord/go_json_benchmark/pkg/utils/JsonType"
	"github.com/NineLord/go_json_benchmark/pkg/utils/Randomizer"
	"log"
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

func (jsonGenerator jsonGenerator) generateFullTree() (map[string]interface{}, error) {
	root := make(map[string]interface{})

	// #region Edge Cases
	if jsonGenerator.depth == 0 {
		return root, nil
	} else if jsonGenerator.depth == 1 {
		jsonGenerator.addLeafChildrenToMap(root)
		return root, nil
	}
	// #endregion

	// #region First Level
	jsonGenerator.addNoneLeafChildrenToMap(root)
	// #endregion

	// #region Middle Levels
	currentNodes := make([]interface{}, 1)
	currentNodes = append(currentNodes, root)
	// nextLevelNodes := make([]interface{}, 0) // Shaked-TODO
	lastLevel := jsonGenerator.depth - 1

	for _level := uint(0); _level < lastLevel; _level++ {
		for 0 < len(currentNodes) {
			var currentNode interface{}
			currentNode, currentNodes = currentNodes[len(currentNodes)-1], currentNodes[:len(currentNodes)-1] // Pop
			currentNodeType := reflect.TypeOf(currentNode).Kind()
			switch currentNodeType {
			case reflect.Slice:
				panic("TODO")
			case reflect.Map:
				for _, nextLevelNode := range currentNode.(map[string]interface{}) {
					nextLevelNodeType := reflect.TypeOf(nextLevelNode).Kind()
					switch nextLevelNodeType {
					case reflect.Slice:
						nextLevelNode = jsonGenerator.addNoneLeafChildrenToSlice(nextLevelNode.([]interface{}))
					case reflect.Map:
					default:
						panic(fmt.Sprintf("generateFullTree unknown JSON type: %d", nextLevelNodeType))
					}
				}
				panic("TODO")
			default:
				panic(fmt.Sprintf("generateFullTree unknown JSON type: %d", currentNodeType))
			}

		}
	}
	// #endregion

	log.Panic("TODO: continue here")

	return root, nil
}

// #region Helper methods
func (jsonGenerator jsonGenerator) getRandomNodeCharacter() *rune {
	index := Randomizer.GetRandomIndexFromArray(jsonGenerator.characterPoll)
	char := jsonGenerator.characterPoll[index]
	return &char
}

func (jsonGenerator jsonGenerator) getRandomNodeName() string {
	var stringBuilder strings.Builder
	for count := uint(0); count < jsonGenerator.numberOfLetters; count++ {
		stringBuilder.WriteRune(*jsonGenerator.getRandomNodeCharacter())
	}
	return stringBuilder.String()
}

func (jsonGenerator jsonGenerator) addNoneLeafChildrenToSlice(array []interface{}) []interface{} {
	for _nodeCount := uint(0); _nodeCount < jsonGenerator.numberOfChildren; _nodeCount++ {
		childNode := JsonType.GetRandomNoneLeafJson(jsonGenerator.numberOfChildren)
		array = append(array, childNode)
	}
	return array
}

func (jsonGenerator jsonGenerator) addNoneLeafChildrenToMap(object map[string]interface{}) {
	for _nodeCount := uint(0); _nodeCount < jsonGenerator.numberOfChildren; _nodeCount++ {
		childNodeName := jsonGenerator.getRandomNodeName()
		childNode := JsonType.GetRandomNoneLeafJson(jsonGenerator.numberOfChildren)
		object[childNodeName] = childNode
	}
}

func (jsonGenerator jsonGenerator) addNoneLeafChildren(nextLevelNode interface{}, nextLevelNodes []interface{}) {
	reflectionType := reflect.TypeOf(nextLevelNode)
	switch reflectionType.Kind() {
	case reflect.Slice:
		nextLevelNode = jsonGenerator.addNoneLeafChildrenToSlice(nextLevelNode.([]interface{}))
	case reflect.Map:
	default:
		panic("TODO")
	}
	panic("TODO")
}

func (jsonGenerator jsonGenerator) addLeafChildrenToMap(object map[string]interface{}) {
	for _nodeCount := uint(0); _nodeCount < jsonGenerator.numberOfChildren; _nodeCount++ {
		childNodeName := jsonGenerator.getRandomNodeName()
		childNode := JsonType.GetRandomLeafJson()
		object[childNodeName] = childNode
	}
}

// #endregion
