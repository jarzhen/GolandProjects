package unittest

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"
)

func TestRegexMatch(t *testing.T) {
	var filePath string = "C:\\Users\\jiazhen\\Desktop\\log转csv\\test.log"
	var regex string = "^((\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{3}) (\\[[\\w\\.-]+\\])(\\[\\w*?\\]) ([A-Z]+) +([\\w\\.]*?) - \\[([\\w<>]+), (\\d+)\\] - ?(.*?))$|^(.*?)$"
	var headTitleInfo string = "日志日期:2@d|线程名:3@s|跟踪标识:4@s|日志级别:5@s|记录器:6@s|方法名:7@s|代码行号:8@n|日志信息:9@j"
	fmt.Println("headTitleInfo = ", headTitleInfo)
	RegexMatch(filePath, regex)
	regReplace()
}

func regReplace() {
	re := regexp.MustCompile("[\u4e00-\u9fa5]+:")
	rep := re.ReplaceAllString("日志日期:2@d|", "")
	fmt.Println("headTitleInfo = ", rep)
}

func RegexMatch(filePath string, regex string) {
	// open the file
	file, err := os.Open(filePath)
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	var i int = 0
	reg1 := regexp.MustCompile(regex)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	// read line by line
	for fileScanner.Scan() {
		i++
		var line string = fileScanner.Text()
		fmt.Println(line)
		fmt.Println(i)
		//根据规则提取关键信息
		result1 := reg1.FindStringSubmatch(line)
		fmt.Println("result = ", result1)
		fmt.Println("result0 = ", result1[0])
		fmt.Println("result2 = ", result1[2])
		fmt.Println("result3 = ", result1[3])
		fmt.Println("result4 = ", result1[4])
		fmt.Println("result5 = ", result1[5])
		fmt.Println("result6 = ", result1[6])
		fmt.Println("result7 = ", result1[7])
		fmt.Println("result8 = ", result1[8])
		fmt.Println("result9 = ", result1[9])
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}
