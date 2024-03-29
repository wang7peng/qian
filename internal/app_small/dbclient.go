package app_small

import (
	"fmt"
	"os/exec"
	SQL "qianDev/consts"
	ConfFile "qianDev/dialplan"
)

// 非root权限连接 sql需要密码
func QianDBA(sql string) string {
	out, _ := exec.Command("bash", "-c", "whoami").Output()
	if string(out[:len(out)-1]) != "root" {
		return fmt.Sprintf("mysql -u asterisk -plike2024 -e \"%s\" ", sql)
	}
	return fmt.Sprintf("mysql -e \"%s\" ", sql)

}

// 便捷的数据显示, 用来显示列, 直接去头, 省去 bash 中管道接 head
func QianDBA_Unit(sql string) string {
	out, _ := exec.Command("bash", "-c", "whoami").Output()
	if string(out[:len(out)-1]) != "root" {
		return fmt.Sprintf("mysql -u asterisk -plike2024 -Ne \"%s\" ", sql)
	}
	return fmt.Sprintf("mysql -Ne \"%s\" ", sql)
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
		fmt.Println("no Ext Number specified!")
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
		id, aors, auth, context := endpoint, endpoint, endpoint, ""
		fmt.Print("上下文(默认 sets): ")
		fmt.Scanln(&context)
		if len(context) == 0 {
			context = "sets"
		}
		cmd := fmt.Sprintf(QianDBA(SQL.E[3]), id, "transport-udp", aors, auth, context, "all", "ulaw", "no")
		_, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("创建分机失败")
			return
		}

		// 在配置 extension.conf 中定义一个提示
		ConfFile.Insert_Hint(context, endpoint, "/etc/asterisk/extensions.conf")
	}
}

/*
功能: 命令行创建分机, 一行创建

  - app -c 2024
  - app -c 2024 like2024 qian-ctx
*/
func CreateExtension(params []string) {
	// 三张表确定一个分机
	endpoint, password, context := params[0], "", ""

	if exist := Check_Aors(endpoint); !exist {
		cmd := fmt.Sprintf(QianDBA(SQL.E[1]), endpoint, 1)
		exec.Command("bash", "-c", cmd).Output()
	}

	if exist := Check_Auths(endpoint); !exist {

		if len(params) > 1 {
			password = params[1]
		} else {
			fmt.Print("设置密码: ")
			fmt.Scanln(&password)
		}

		id := endpoint
		cmd := fmt.Sprintf(QianDBA(SQL.E[2]), id, "userpass", password, endpoint)
		exec.Command("bash", "-c", cmd).Output()
	}

	// 检测是否已经有该设备
	if exist := Check_Endpoint(endpoint); !exist {
		id, aors, auth := endpoint, endpoint, endpoint
		if len(params) > 2 {
			context = params[2]
		} else {
			fmt.Print("上下文(默认 sets): ")
			fmt.Scanln(&context)
		}
		if len(context) == 0 {
			context = "sets"
		}

		cmd := fmt.Sprintf(QianDBA(SQL.E[3]), id, "transport-udp", aors, auth, context, "all", "ulaw", "no")
		_, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("创建分机失败")
			return
		}

		// 在配置 extension.conf 中定义一个提示
		ConfFile.Insert_Hint(context, endpoint, "/etc/asterisk/extensions.conf")
		fmt.Println("分机 " + endpoint + " 创建完成!")
	}
}

// 功能: 根据提供的分机号删除对应的记录
func DeleteExtension(endpoint string) {

	if !Check_Aors(endpoint) || !Check_Auths(endpoint) {
		fmt.Println("无需删除, 系统没有该分机 " + endpoint)
		return
	}

	// 删除配置文件中与分机相配的 endpoint,hint,xxx 一行
	ConfFile.Delete_Hint(endpoint, "/etc/asterisk/extensions.conf")

	cmd := fmt.Sprintf(SQL.E[201]+SQL.E[202]+SQL.E[203], endpoint, endpoint, endpoint)
	_, err := exec.Command("bash", "-c", QianDBA(cmd)).Output()
	if err != nil {
		fmt.Printf("删除分机 %s 失败", endpoint)
		return
	}
	fmt.Println("分机 " + endpoint + " 已经删除!")
}

// 功能: 通过单个分机名字获取其上下文
func GetContext(id string) string {
	cmd := fmt.Sprintf(QianDBA_Unit(SQL.E[111]), id)
	out, _ := exec.Command("bash", "-c", cmd).Output()
	if len(out) == 0 {
		return ""
	}
	return string(out[:len(out)-1])
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
	// 查到数据返回 1, 不存在返回 0
	return (string(out) == "1")
}
