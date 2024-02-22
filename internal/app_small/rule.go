package app_small

import (
	"fmt"
	"strings"

	ConfFile "qianDev/dialplan"
)

func ObeyRule(ruleStr string) {
	if len(ruleStr) == 0 || len(ruleStr) < 7 {
		return
	}

	// A ?--?--? B
	ruleStr, ok := CanRegular(ruleStr)
	if !ok {
		ShowWarn()
		return
	}

	// 规范后的字符串一定可以通过空格被切成三部分
	strArr := strings.Split(ruleStr, " ")

	// 两侧空格已经预防掉, 可以直接用, 不要再 Trim
	A, B := strArr[0], strArr[2]

	switch strArr[1] {
	case "---->", "---1-":
		fmt.Println(A, " can call ", B)
		ConfFile.Insert_Dial(GetContext(A), B, 30, "/etc/asterisk/extensions.conf")

	case "<----", "-1---":
		ConfFile.Insert_Dial(GetContext(B), A, 30, "/etc/asterisk/extensions.conf")

	case "-----", "<--->", "-1-1-":
		ConfFile.Insert_Dial(GetContext(A), B, 30, "/etc/asterisk/extensions.conf")
		ConfFile.Insert_Dial(GetContext(B), A, 30, "/etc/asterisk/extensions.conf")

	case "--x--", "--0--":
		// TODO

	default:
		fmt.Println("更多格式应用 敬请期待...")
	}

}

/*
格式规定:

	A-----B     A ----- B     A --1-- B
	A---->B     A ----> B     A ---1- B
	A<----B     A <---- B     A -1--- B
	A<--->B     A <---> B     A -1-1- B
	A--x--B     A --x-- B     A --0-- B

实现原理:

	先除去所有空格, 再在中间组合标志前后加一个空
*/
func CanRegular(str string) (string, bool) {
	goodStr := ""
	defer func() {
		fmt.Println(goodStr)
	}()

	if len(str) == 0 || len(str) < 7 {
		return str, false
	}

	if strings.Contains(str, " ") {
		str = strings.Replace(str, " ", "", -1) // 清光空格
	} else {
		goodStr = str
	}

	// 两个检索不能并列定义, 因为首位置插入后, 尾位置会对不上号
	fInt := strings.Index(str, "-")
	if fInt == -1 {
		return str, false // 等效 !strings.Contains(str, "-")
	}

	// 在左侧 < 前面插入空格, 不需要前移 -2
	if str[fInt-1] == '<' {
		str = str[:fInt-1] + " " + str[fInt-1:]
	} else {
		str = str[:fInt] + " " + str[fInt:] // 直接插在第一个标志 - 上
	}

	// 在 > 后面插入, 需要后移 +2
	if eInt := strings.LastIndex(str, "-"); str[eInt+1] == '>' {
		goodStr = str[:eInt+2] + " " + str[eInt+2:]
	} else {
		goodStr = str[:eInt+1] + " " + str[eInt+1:]
	}

	return goodStr, true
}
