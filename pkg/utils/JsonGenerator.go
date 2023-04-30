package utils

func GenerateJson(characterPoll string, numberOfLetters uint, depth uint, numberOfChildren uint) (*map[string]interface{}, error) {
	var data map[string]interface{} = map[string]interface{}{"apple": 5}
	return &data, nil
}
