package JsonGenerator

import (
	"github.com/NineLord/go_json_benchmark/pkg/utils/Randomizer"
	"log"
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

func GenerateJson(characterPoll string, numberOfLetters uint, depth uint, numberOfChildren uint) (*map[string]interface{}, error) {
	return newJsonGenerator(characterPoll, numberOfLetters, depth, numberOfChildren).generateFullTree()
}

func (jsonGenerator jsonGenerator) generateFullTree() (*map[string]interface{}, error) {
	root := map[string]interface{}{}

	//#region Edge Cases
	if jsonGenerator.depth == 0 {
		return &root, nil
	} else if jsonGenerator.depth == 1 {
		//jsonGenerator.addLeafChildrenToMap(&root) // Shaked-TODO
		return &root, nil
	}
	//#endregion

	log.Panic("TODO: continue here")

	return &root, nil
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

func addNodeLeafChildrenToArray() {
	log.Panic("TODO")
}

//#endregion
