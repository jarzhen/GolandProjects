package unittest

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	var filePath = "C:\\Users\\jiazhen\\Desktop\\log转csv\\test.log"
	ReadFile(filePath)
}

func ReadFile(filePath string) {
	// open the file
	file, err := os.Open(filePath)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}
