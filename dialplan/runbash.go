/* 使用执行bash命令的方式操作配置文件
 */
package dialplan

import (
	"os/exec"

	"fmt"
	"strconv"
)

// 执行 wc -l 获取文件总行数
func getFileRows(file string) int {
	cmd := fmt.Sprintf("cat %s | wc --line ", file)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return -1
	}
	valueStr := string(out[:len(out)-1])
	count, _ := strconv.Atoi(valueStr)
	return count
}

// 执行 sed -n 获取文件指定行的内容
func getLineInFile(row int, file string) string {
	cmd := fmt.Sprintf("sed -n %dp %s", row, file)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return ""
	}
	if len(out) == 0 {
		return ""
	}

	valueStr := string(out[:len(out)-1])
	return valueStr
}

/*
功能: 执行 sed -n 获取文件指定两个行区间的内容

	e.g sed -n '20, 30p' /etc/.../extensions.conf
*/
func getLinesInRowRange(rowStart, rowEnd int, file string) string {
	cmd := fmt.Sprintf("sed -n '%d, %dp' %s", rowStart, rowEnd, file)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return ""
	}
	if len(out) == 0 {
		return ""
	}

	valueStr := string(out[:len(out)-1])
	return valueStr
}

/*
功能: 执行 sed -i 插入单行内容到文件指定行

	e.g sed -i 'ni\xxx' /etc/xxx/extensions.conf
*/
func InsertLineToFile(n int, row, file string) error {
	cmd := fmt.Sprintf("sudo sed -i '%di\\%s' %s ", n, row, file)

	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err
	}
	return nil
}

// 执行 grep xxx | wc 判断文字串中有没有指定的格式串
func checkStringInText(text, sonStr string) bool {
	cmd := fmt.Sprintf("echo \"%s\" | grep %s | wc -l", text, sonStr)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return false
	}

	valueStr := string(out[:len(out)-1])
	count, _ := strconv.Atoi(valueStr)

	return count > 0
}
