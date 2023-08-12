package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const filename string = "./fav_number.json"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your favorite number: ")
	scanner.Scan()

	number, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println("That is not a valid number!!")
		return
	}

	saveNumber(number)
	readSavedNumber(filename)

}

func readSavedNumber(filename string) {
	jsonFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File does not exists!!")
		panic(err)
	}

	jsonValue := map[string]interface{}{}
	json.Unmarshal(jsonFile, &jsonValue) // convert the JSON to valid map type

	fmt.Printf("I know your favorite number!!. It's %v\n", jsonValue["number"])
}

func saveNumber(number int64) {
	// Convert the number to JSON format
	jsonBytes, _ := json.Marshal(map[string]interface{}{"number": number})

	// Create a JSON file if not exists,
	jsonFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close() // close the file at the end

	// Save the bytes to a JSON file format
	jsonFile.Write(jsonBytes)
}
