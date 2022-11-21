package unittest

import (
	"fmt"
	"strconv"
	"testing"
)

func TestArray(t *testing.T) {
	arrFunc()
}

func arrFunc() {
	var csvDataLine = []string{strconv.Itoa(1), "否"}
	tmp := append(csvDataLine, "new")
	fmt.Println("csvDataLine = ", csvDataLine)
	fmt.Println("tmp = ", tmp)
	csvDataLine = append(csvDataLine, "new")
	fmt.Println("csvDataLine = ", csvDataLine)
	csvDataLine[1] = "是"
	fmt.Println("csvDataLine = ", csvDataLine)
}
