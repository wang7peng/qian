package main

import (
	"fmt"
	"os/exec"
	SQL "qianDev/consts"
)

// 非root权限连接 sql需要密码
func QianDBA(sql string) string {
	out, _ := exec.Command("bash", "-c", "whoami").Output()
	if string(out[:len(out)-1]) != "root" {
		return fmt.Sprintf("mysql -u asterisk -plike2024 -e \"%s\" ", sql)
	}
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

	if exist := Check_Aors(endpoint); !exist {
		cmd := fmt.Sprintf(QianDBA(SQL.E[1]), endpoint, 1)
		exec.Command("bash", "-c", cmd).Output()
	}

	if exist := Check_Auths(endpoint); !exist {
		fmt.Print("密码: ")
		password := ""
		fmt.Scanln(&password)
		id := endpoint
		cmd := fmt.Sprintf(QianDBA(SQL.E[2]), id, "userpass", password, endpoint)
		exec.Command("bash", "-c", cmd).Output()
	}

	// 检测是否已经有该设备
	if exist := Check_Endpoint(endpoint); !exist {
		id, aors, auth := endpoint, endpoint, endpoint
		cmd := fmt.Sprintf(QianDBA(SQL.E[3]), id, "transport-udp", aors, auth, "sets", "all", "ulaw", "no")
		_, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("创建分机失败")
			return
		}

	}
}

func CreateExtension(endpoint string) {
	// 三张表确定一个分机

	if exist := Check_Aors(endpoint); !exist {
		cmd := fmt.Sprintf(QianDBA(SQL.E[1]), endpoint, 1)
		exec.Command("bash", "-c", cmd).Output()
	}

	if exist := Check_Auths(endpoint); !exist {
		fmt.Print("密码: ")
		password := ""
		fmt.Scanln(&password)
		id := endpoint
		cmd := fmt.Sprintf(QianDBA(SQL.E[2]), id, "userpass", password, endpoint)
		exec.Command("bash", "-c", cmd).Output()
	}

	// 检测是否已经有该设备
	if exist := Check_Endpoint(endpoint); !exist {
		id, aors, auth := endpoint, endpoint, endpoint
		cmd := fmt.Sprintf(QianDBA(SQL.E[3]), id, "transport-udp", aors, auth, "sets", "all", "ulaw", "no")
		_, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("创建分机失败")
			return
		}
		fmt.Println("分机 " + endpoint + " 创建完成!")
	}
}

// 功能: 检测是否存在记录地址
func Check_Aors(id string) bool {
	cmd := fmt.Sprintf(QianDBA(SQL.E[101])+"|sed -n '2p' ", id)
	out, _ := exec.Command("bash", "-c", cmd).Output()
	return (string(out[:len(out)-1]) == "1")
}

// 功能: 检测是否存在身份认证
func Check_Auths(id string) bool {
	cmd := fmt.Sprintf(QianDBA(SQL.E[102])+"|sed -n '2p' ", id)
	out, _ := exec.Command("bash", "-c", cmd).Output()
	return (string(out[:len(out)-1]) == "1")
}

// 功能: 检测是否存在指定的终端
func Check_Endpoint(endpoint string) bool {
	cmd := fmt.Sprintf(QianDBA(SQL.E[103])+"|sed -n '2p' ", endpoint)
	out, _ := exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]
	fmt.Println(string(out))
	return (string(out) == "1")
}