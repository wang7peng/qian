package dialplan

/* 配置文件在的 dialplan 修改 */

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ObjContext struct {
	line int
	name string
}

func RowStatement_Dial(id string, timeout int) string {
	// e.g exten => id,1,Dial(PJSIP/id)
	// e.g exten => id,1,Dial(PJSIP/id,30)
	return fmt.Sprintf("exten => %s,1,Dial(PJSIP/%s, %d)",
		id, id, timeout)
}

func RowStatement_Hint(id string) string {
	// e.g exten => id,hint,PJSIP/id&Custom:DNDid,CustomPresence:id
	return fmt.Sprintf("exten => %s,hint,PJSIP/%s&Custom:DND%s,CustomPresence:%s",
		id, id, id, id)
}

/*
功能: 在 extension.conf 中写定义一个和指定分机相关的提示

  - 如果没有上下文或者上下文在最后，直接在末尾插入;
  - 如果上下文在中间某行, 就利用 sed 命令插入
*/
func Insert_Hint(context, endpoint, filepath_conf string) {
	rowContext := "[" + strings.TrimSpace(context) + "]"

	// 获取文件中所有 [context] 所在的行
	contextArr := getContext(filepath_conf)

	// 本次上下文及其后面的上下文
	exist, pos := false, 0
	for i := 0; i < len(contextArr); i++ {
		if contextArr[i].name == rowContext {
			exist, pos = true, i
		}
	}

	if !exist {
		/*
			[context]
			exten => id,hint,PJSIP/id&Custom:DNDid,CustomPresence:id
		*/
		content := rowContext + "\n" +
			RowStatement_Hint(endpoint) + "\n"

		appendToFile(filepath_conf, content)

	} else if pos == len(contextArr)-1 {
		content := RowStatement_Hint(endpoint) + "\n"
		appendToFile(filepath_conf, content)

	} else {
		// 将原先这一行的内容顶掉, 不用插在前一行
		lineInsert := contextArr[pos+1].line

		content := RowStatement_Hint(endpoint)

		err := InsertLineToFile(lineInsert, content, filepath_conf)
		if err != nil {
			fmt.Println("插入失败")
			return
		}

		fmt.Printf("在 extensions.conf 的 %d 行插入hint成功! \n", lineInsert)
	}

}

/*
功能: 在 extensions.conf 中代表指定分机提示的一行去除

  - 不需要找要删除的分机的context, 每个分机的 hint 语句有且只有一行
*/
func Delete_Hint(endpoint, filepath_conf string) {

	numArr := getRowsNum(endpoint+",hint", filepath_conf)
	if numArr[0] == 0 {
		return // 没有找到无需删除
	}

	// 只管 numArr[0] , 就是要删除的数据
	// numArr 其他不管, 因为删除后, 文本后部分都上提一行, 和实际已经对不上了
	err := deleteLineInFile(numArr[0], filepath_conf)
	if err != nil {
		fmt.Println("删除失败: ", err.Error())
	}
}

/*
功能: 在 extensions.conf 中指定的上下文里定义一个能打通分机的连接

  - A ----> B 在A的上下文中连接B
*/
func Insert_Dial(context, endpoint string, timeout int, filepath_conf string) {
	contextInConf := "[" + context + "]"
	contextArr, rows := getContext(filepath_conf), getFileHeight(filepath_conf)

	// 优先考虑末尾, 可能需要找的上下文就在末尾
	objLast := contextArr[len(contextArr)-1]
	if objLast.name == contextInConf {
		fmt.Printf("在行区间 [%d, %d] 插入... ", objLast.line+1, rows)
		// sed -n 'a, bp' xxx.conf | grep xxx
		content := getLinesInRowRange(objLast.line+1, rows, filepath_conf)

		/*
			[context]
			exten => id,1,Dial(PJSIP/id,30)
		*/
		if !checkStringInText(content, endpoint+",1,Dial") {
			rowNumber, content := objLast.line+1, RowStatement_Dial(endpoint, 30)
			err := InsertLineToFile(rowNumber, content, filepath_conf)
			if err != nil {
				fmt.Println("插入失败: ", err.Error())
			} else {
				fmt.Println("ok")
			}
		}
		return
	}

	i := 0
	for ; i < len(contextArr); i++ {
		if contextArr[i].name == contextInConf {
			objCurr, objNext := contextArr[i], contextArr[i+1]
			fmt.Printf("在行区间 [%d, %d] 插入... ", objCurr.line+1, objNext.line-1)

			//  sed -n 'a, bp' xxx.conf | grep xxx
			content := getLinesInRowRange(objCurr.line+1, objNext.line-1, filepath_conf)
			if !checkStringInText(content, endpoint+",1,Dial") {
				rowNumber, content := objCurr.line+1, RowStatement_Dial(endpoint, 30)
				err := InsertLineToFile(rowNumber, content, filepath_conf)
				if err != nil {
					fmt.Println("插入失败: ", err.Error())
				} else {
					fmt.Println("ok")
				}
			}

			break // 只找第一个
		}
	}

	if i == len(contextArr) {
		fmt.Println("分机未通过写 dialplan 方式实现, 此路不通...")
		return
	}
}

// 获取配置文件的总行数
func getFileHeight(file string) int {
	fi, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return -1
	}
	defer fi.Close()

	fd := bufio.NewReader(fi)
	count := 0
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}
	return count
}

// 所有 [ 起头的行
func getContext(file string) []ObjContext {
	contextArr := make([]ObjContext, 0)

	fi, _ := os.Open(file)
	r := bufio.NewReader(fi)
	for i := 1; ; i++ {
		lineBytes, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if lineBytes[0] == '[' {
			lineBytes = lineBytes[:len(lineBytes)-1] // 去除末尾 \n
			contextArr = append(contextArr, ObjContext{line: i, name: string(lineBytes)})
		}
	}

	return contextArr
}

// ref: https://blog.csdn.net/choumin/article/details/90319294
func appendToFile(absPath, str string) {
	file, err := os.OpenFile(absPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error %s!\n", err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString(str)
	if err != nil {
		fmt.Println("Err: ", err.Error())
		return
	}

	fmt.Printf("在 %s 的末尾插入hint成功! \n", filepath.Base(absPath))
}

func deleteLineInFile(row int, file string) error {
	f, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(f), "\n")

	// 创建一个新的内容，不包含要删除的行
	var newContent []string
	for i, line := range lines {
		if i != row-1 { // 行号减1是因为slice的索引是从0开始的
			newContent = append(newContent, line)
		}
	}

	// 将新的内容写回文件
	return os.WriteFile(file, []byte(strings.Join(newContent, "\n")), 0644)
}
