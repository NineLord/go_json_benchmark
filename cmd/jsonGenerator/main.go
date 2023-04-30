package main

import (
	"fmt"
	"github.com/NineLord/go_json_benchmark/pkg/utils"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli/v2"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

func main() {
	app := &cli.App{
		Name:      "jsonGenerator",
		Usage:     "Generates JSON file for testing",
		ArgsUsage: "[pathToSaveFile]", // Can't add description to arguments with this package
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:    "numberOfLetters",
				Aliases: []string{"n"},
				Usage:   "The total number of letters that each generated node name will have",
				Value:   7,
			},
			&cli.UintFlag{
				Name:    "depth",
				Aliases: []string{"d"},
				Usage:   "The depth of the JSON tree",
				Value:   100,
			},
			&cli.UintFlag{
				Name:    "numberOfChildren",
				Aliases: []string{"m"},
				Usage:   "The number of children each node should have",
				Value:   6,
			},
			&cli.BoolFlag{
				Name:    "print",
				Aliases: []string{"P"},
				Usage:   "Print the resulting JSON instead of saving it to a file",
				Value:   false,
			},
		},
		Action: cliAction,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func parsePathToSaveFile(arguments *cli.Args) string {
	result := (*arguments).Get(0)
	if result == "" {
		if homeDir, err := os.UserHomeDir(); err == nil {
			result = homeDir
		} else if executable, err := os.Executable(); err == nil {
			result = filepath.Dir(executable)
		} else {
			log.Fatal("Didn't get result and couldn't get default path, error: ", err)
		}
	}
	result += "/generatedJson.json"
	return result
}

func cliAction(arguments *cli.Context) error {
	args := arguments.Args()
	pathToSaveFile := parsePathToSaveFile(&args)

	numberOfLetters := arguments.Uint("numberOfLetters")
	depth := arguments.Uint("depth")
	numberOfChildren := arguments.Uint("numberOfChildren")
	printFlag := arguments.Bool("print")
	if json, err := utils.GenerateJson(ALPHABET, numberOfLetters, depth, numberOfChildren); err == nil {
		if printFlag {
			fmt.Println(json)
			return nil
		} else {
			err := os.WriteFile(pathToSaveFile, []byte(strconv.Itoa(json)), 0644)
			fmt.Println("JSON was saved to", pathToSaveFile)
			return err
		}
	} else {
		return err
	}
}
