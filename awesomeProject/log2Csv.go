package main

import (
	"bufio"
	"container/list"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	fmt.Println(args)
	var filePath string = args[1]
	var regex string = args[2]
	var headTitleInfo string = args[3]
	fmt.Println("headTitleInfo = ", headTitleInfo)
	Log2Csv(filePath, regex, headTitleInfo)
}

func replaceAndSplit(metaInfo string, regex string, replacement string, split string) []string {
	// 正则替换
	re, _ := regexp.Compile(regex)
	rep := re.ReplaceAllString(metaInfo, replacement)
	headList := strings.Split(rep, split)
	return headList
}

func processLogByType(typ string, logString string) string {
	// d:date,日期
	// n:mumber,数字
	// j:json, json
	// s:string,字符

	var result string = logString
	switch typ {
	case "d":
		result = fmt.Sprintf("=\"%s\"", logString)
	case "n":
		result = fmt.Sprintf("%s", logString)
	case "j":
		result = fmt.Sprintf("%s", strings.Replace(logString, "\"", "\"\"", -1))
	case "s":
		result = fmt.Sprintf("%s", logString)
	default:
		result = fmt.Sprintf("%s", logString)
	}
	return result
}

func Log2Csv(filePath string, regex string, metaInfo string) {

	// open the file
	file, err := os.Open(filePath)
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	headList := replaceAndSplit(metaInfo, ":\\d@\\w", "", "|")
	var arr1 = []string{"日志行号", "是否零散日志"}
	headInfo := append(arr1, headList...)
	fmt.Printf("%v", headInfo)
	csvData := list.New()
	csvData.PushBack(headInfo)
	var lastMatchingInfo []string = strings.SplitN(strings.Repeat(",", len(headList)-1), ",", len(headList))
	fmt.Printf("%v", lastMatchingInfo)
	reg1 := regexp.MustCompile(regex)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	// read line by line
	var i int = 0
	var j int = 0
	var k int = 0
	for fileScanner.Scan() {
		i++
		var line string = fileScanner.Text()
		fmt.Println(line)
		fmt.Println(i)
		//根据规则提取关键信息
		result := reg1.FindStringSubmatch(line)
		if result[1] != "" {
			j++
			fmt.Println("result = ", result[1])
			regMetaData := replaceAndSplit(metaInfo, "[\u4E00-\u9FA5]+:", "", "|")
			var csvDataLine = []string{strconv.Itoa(i), "否"}
			for index, value := range regMetaData {
				metaArr := strings.Split(value, "@")
				atoi, err := strconv.Atoi(metaArr[0])
				if err != nil {
					fmt.Println("Atoi err")
					return
				}
				cellValue := processLogByType(metaArr[1], result[atoi])
				csvDataLine = append(csvDataLine, cellValue)
				lastMatchingInfo[index] = cellValue
			}
			csvData.PushBack(csvDataLine)
		} else {
			k++
			fmt.Println("result = ", result[10])
			copyArr := lastMatchingInfo
			copyArr[len(lastMatchingInfo)-1] = processLogByType("j", result[0])
			var csvDataLine = []string{strconv.Itoa(i), "是"}
			csvDataLine = append(csvDataLine, copyArr...)
			csvData.PushBack(csvDataLine)
		}
	}
	fmt.Printf("i=%d;j=%d;k=%d", i, j, k)
	fmt.Printf("i==j+k?%t", i == (j+k))
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
	csvFilePath := strings.Replace(filePath, ".log", ".csv", 1)
	writeCSV(csvFilePath, csvData)
}

// csv文件写入
func writeCSV(path string, data *list.List) {

	//OpenFile读取文件，不存在时则创建，使用追加模式
	File, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Println("文件打开失败！")
	}
	defer File.Close()

	//创建写入接口
	WriterCsv := csv.NewWriter(File)
	for e := data.Front(); e != nil; e = e.Next() {
		//写入一条数据，传入数据为切片(追加模式)
		str := e.Value.([]string)
		err1 := WriterCsv.Write(str)
		if err1 != nil {
			log.Println("WriterCsv写入文件失败")
		}
	}

	WriterCsv.Flush() //刷新，不刷新是无法写入的
	log.Println("数据写入成功...")
}
