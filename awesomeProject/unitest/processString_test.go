package unittest

import (
	"fmt"
	"strings"
	"testing"
)

func TestProcessString(t *testing.T) {
	var headTitleInfo = "日志日期:2@d|线程名:3@s|跟踪标识:4@s|日志级别:5@s|记录器:6@s|方法名:7@s|代码行号:8@n|日志信息:9@j"
	ProcessString(headTitleInfo)
}

func ProcessString(metaInfo string) {
	//// 正则替换
	//re, _ := regexp.Compile(":\\d@\\w")
	//rep := re.ReplaceAllString(metaInfo, "")
	//headList := strings.Split(rep, "|")
	//var arr1 = []string{"日志行号", "是否零散日志"}
	//headInfo := append(arr1, headList...)
	//fmt.Printf("%v", headInfo)
	//
	//csvData := list.New()
	//csvData.PushBack(headInfo)

	splitN()
}

func splitN() {
	n := strings.SplitN(strings.Repeat(",", 7), ",", 8)
	fmt.Printf("%v", n)
}
