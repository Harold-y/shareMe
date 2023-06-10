package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	AdminToken       string
	Delim            string
	UploadFolder     string
	MaxUploadSize    int
	DestroyOnStart   string
	DailyDestroyHour int
	Port             string
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
			if key == "Delim" {
				Delim = value
			}
			if key == "UploadFolder" {
				UploadFolder = value
			}
			if key == "MaxUploadSize" {
				MaxUploadSize, err = strconv.Atoi(value)
				if err != nil {
					fmt.Println("MaxUploadSize Error")
				}
			}
			if key == "DestroyOnStart" {
				DestroyOnStart = value
			}
			if key == "DailyDestroyHour" {
				DailyDestroyHour, err = strconv.Atoi(value)
				if err != nil || DailyDestroyHour > 25 || DailyDestroyHour < 0 {
					fmt.Println("Destroy Hour Error")
				}
			}
			if key == "Port" {
				Port = value
			}
		}
	}
}
