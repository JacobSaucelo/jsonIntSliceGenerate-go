package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

type dataType struct {
	Data []int `json:"data"`
}

var folderName = "data"

func main() {
	value := generateData()

	fPath := filepath.Join(folderName, "data.json")
	err := saveFile(fPath, value)
	if err != nil {
		fmt.Println("error saving file")
		return
	}
}

func generateData() dataType {
	var data = []int{}
	for i := 0; i < 1000; i++ {
		data = append(data, rand.Intn(1000-1)+1)
	}

	return dataType{
		Data: data,
	}

}

func saveFile(fpath string, data dataType) error {
	saveFile, err := os.Create(fpath)
	if err != nil {
		fmt.Println("error opening save file", saveFile)
		return err
	}
	defer saveFile.Close()

	fmt.Println("json", data)
	encoder := json.NewEncoder(saveFile)
	err = encoder.Encode(data)
	if err != nil {
		fmt.Println("error encoding data to savefile")
		return err
	}

	fmt.Println("successfully generated data.json")
	return nil

}
