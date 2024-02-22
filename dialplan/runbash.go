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
