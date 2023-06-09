package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	AdminToken string
)

func SetConfig() {
	// Open the configuration
	configPath := "config.txt"
	configFile, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) >= 2 {
			key, value := words[0], words[1]
			if key == "AdminToken" {
				AdminToken = value
			}
		}
	}
}
