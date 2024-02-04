package dialplan

/* 配置文件在的 dialplan 修改 */

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type ObjContext struct {
	line int
	name string
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
		// sed -n 'ni\xxx' /a/b.conf
		cmd := fmt.Sprintf("sudo sed -i '%di\\%s' %s ", lineInsert, content, filepath_conf)

		_, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("插入失败")
			return
		}

		fmt.Printf("在 extensions.conf 的 %d 行插入hint成功! \n", lineInsert)
	}

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
