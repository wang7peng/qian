package main

import (
	"fmt"
	"os/exec"
	SQL "qianDev/consts"
)

func QianDBA(sql string) string {
	return fmt.Sprintf("mysql -e \"%s\" ", sql)
}

/**
 * 功能: 创建话机终端, [1] P48
 *
 * 往 ps_aors ps_auths ps_endpoints 各插入一条记录
 */
func CreateEndPoint() {
	fmt.Print("Endpoint Name: ")
	endpoint := "" // "0002024A1"
	fmt.Scanln(&endpoint)

	if len(endpoint) == 0 {
		fmt.Println("输入不行")
		Usage()
		return
	}

	fmt.Println("开始对PJSIP 终端", endpoint, "设置")

	cmd := fmt.Sprintf(QianDBA(SQL.E[1]), "0002024A1")

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))

	cmd = fmt.Sprintf(QianDBA(SQL.E[2]), "0002024A1", "userpass", "not very secure", "0002024A1")
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	out = out[:len(out)-1]
	fmt.Println(string(out))

	cmd = fmt.Sprintf(QianDBA(SQL.E[3]), "0002024A1", "transport-udp", "0002024A1", "0002024A1", "sets", "all", "ulaw", "no")
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
		return
	}
	out = out[:len(out)-1]
	fmt.Println(string(out))
}
