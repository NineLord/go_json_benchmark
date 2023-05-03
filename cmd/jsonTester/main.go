package main

// #region Imports
import (
	"fmt"
	"github.com/NineLord/go_json_benchmark/pkg/utils/JsonGenerator"
	jsoniter "github.com/json-iterator/go"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// #endregion

var CharacterPoll = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz!@#$%&"

func getDefaultPathToSaveFile() cli.Path {
	if homeDir, err := os.UserHomeDir(); err == nil {
		return homeDir
	} else if executable, err := os.Executable(); err == nil {
		return filepath.Dir(executable)
	} else {
		panic(fmt.Sprintf("Didn't get result and couldn't get default path, error: %s", err))
	}
}

func main() {
	app := &cli.App{
		Name:      "jsonTester",
		Usage:     "Tests JSON manipulations",
		ArgsUsage: "<jsonPath> [testCounter]", // Can't add description to arguments with this package
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "saveFile",
				Aliases: []string{"s"},
				Usage:   "Absolute path to save the excel report file to",
				Value:   getDefaultPathToSaveFile(),
			},
			&cli.UintFlag{
				Name:    "sampleInterval",
				Aliases: []string{"i"},
				Usage:   "The interval in which it will sample the CPU/RAM usage of the system while running the tests, units are in milliseconds",
				Value:   50,
			},
			&cli.UintFlag{
				Name:    "numberOfLetters",
				Aliases: []string{"n"},
				Usage:   "The total number of letters that each generated node name will have in the generated JSON tree",
				Value:   32,
			},
			&cli.UintFlag{
				Name:    "depth",
				Aliases: []string{"d"},
				Usage:   "The depth of the generated JSON tree",
				Value:   255,
			},
			&cli.UintFlag{
				Name:    "numberOfChildren",
				Aliases: []string{"m"},
				Usage:   "The number of children each node should have in the generated JSON tree",
				Value:   16,
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"D"},
				Usage:   "Prints additional debug information",
				Value:   false,
			},
		},
		Action: cliAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func parseArgument(arguments *cli.Args) (cli.Path, uint) {
	jsonPath := (*arguments).Get(0)
	if jsonPath == "" {
		panic("First argument must be the absolute path to the JSON file that will be tested")
	}
	if testCounter := (*arguments).Get(1); testCounter == "" {
		return jsonPath, 5
	} else if testCounter, err := strconv.Atoi(testCounter); err != nil {
		panic(fmt.Sprintf("Second argument should be the number of times will run the tests: %s", err))
	} else {
		return jsonPath, uint(testCounter)
	}
}

func cliAction(arguments *cli.Context) (err error) {
	args := arguments.Args()
	jsonPath, testCounter := parseArgument(&args)
	pathToSaveFile := arguments.Path("saveFile")
	sampleInterval := arguments.Uint("sampleInterval")
	numberOfLetters := arguments.Uint("numberOfLetters")
	depth := arguments.Uint("depth")
	numberOfChildren := arguments.Uint("numberOfChildren")
	debug := arguments.Bool("debug")

	if debug {
		fmt.Println(
			"Arguments:\n",
			"jsonPath:", jsonPath, "\n",
			"testCounter:", testCounter, "\n",
			"pathToSaveFile:", pathToSaveFile, "\n",
			"sampleInterval:", sampleInterval, "\n",
			"numberOfLetters:", numberOfLetters, "\n",
			"depth:", depth, "\n",
			"numberOfChildren:", numberOfChildren, "\n",
			"debug:", debug,
		)
	}

	// Shaked-TODO: create ExcelGenerator here
	// var inputJsonFile string
	var buffer []byte
	if buffer, err = os.ReadFile(jsonPath); err != nil {
		return
	}
	// inputJsonFile = string(buffer)
	/*var inputJsonFile *os.File

	inputJsonFile, err = os.Open(jsonPath)
	if err != nil {
		return
	}
	defer func() {
		closeError := inputJsonFile.Close()
		if err == nil {
			err = closeError
		}
	}()*/
	// valueToSearch := int64(2_000_000_000)

	for count := uint(0); count < testCounter; count++ {
		// Shaked-TODO: create report class to collect Usage and data

		// #region Testing

		// #region Test Generating JSON

		if _, err = JsonGenerator.GenerateJson(CharacterPoll, numberOfLetters, depth, numberOfChildren); err != nil {
			return
		}

		// #endregion

		// #region Test Deserialize JSON

		var inputJsonFile map[string]interface{}
		if err = json.Unmarshal(buffer, &inputJsonFile); err != nil {
			return
		}

		// #endregion

		// #region Test Iterate Iteratively
		// Shaked-TODO
		// #endregion

		// #region Test Iterate Recursively
		// Shaked-TODO
		// #endregion

		// #region Test Serialize JSON

		var buff []byte
		if buff, err = json.Marshal(inputJsonFile); err != nil {
			return
		}
		_ = string(buff)

		// #endregion

		// #endregion
	}

	return nil
}
