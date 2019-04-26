package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	fmt.Println("Start")
	file := getPathCurrent("config.json")
	//config, err := LoadConfiguration(file)
	//if err != nil {
	//fmt.Println(err)
	//}
	//fmt.Println(config)

	change := watchFile(file)
	for b := range change {
		if b {
			config, err := LoadConfiguration(file)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(config)
		}
	}
}

func watchFile(filePath string) <-chan bool {
	initialStat, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}
	change := make(chan bool)
	go func() {
		for {
			stat, err := os.Stat(filePath)
			if err != nil {
				fmt.Println(err)
			}

			if stat.Size() != initialStat.Size() ||
				stat.ModTime() != initialStat.ModTime() {

				change <- true

				initialStat, err = os.Stat(filePath)
				if err != nil {
					fmt.Println(err)
				}
			}

			time.Sleep(1 * time.Second)
		}
	}()
	return change
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)

	err = jsonParser.Decode(&config)
	return config, err
}

func getPathCurrent(nameFile string) string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return filepath.Join(dir, nameFile)
}
