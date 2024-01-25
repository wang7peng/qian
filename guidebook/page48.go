package guidebook

import (
	"fmt"
	"os/exec"

	SQL "qianDev/consts"
)

func QianDBA(sql string) string {
	return fmt.Sprintf("mysql -e \"%s\" ", sql)
}

func CreateEndPointA() {
	endpoint := "0002024A1"
	fmt.Println("开始对PJSIP 终端", endpoint, "设置")

	cmd := fmt.Sprintf(QianDBA(SQL.E[101])+"|sed -n '2p' ", endpoint)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	out = out[:len(out)-1]

	if string(out) == "1" {
		fmt.Printf("record of %s have existed!\n", endpoint)
		return
	}

	cmd = fmt.Sprintf(QianDBA(SQL.E[1]), endpoint, 1)

	out, _ = exec.Command("bash", "-c", cmd).Output()
	fmt.Println(string(out))

	cmd = fmt.Sprintf(QianDBA(SQL.E[2]),
		endpoint, "userpass", "not very secure", "0002024A1")
	out, _ = exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]
	fmt.Println(string(out))

	aors, auth := endpoint, endpoint
	cmd = fmt.Sprintf(QianDBA(SQL.E[3]),
		endpoint, "transport-udp", aors, auth, "sets", "all", "ulaw", "no")
	out, _ = exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]
	fmt.Println(string(out))
}

func CreateEndPointB() {
	endpoint := "0002024B1"
	fmt.Println("开始对PJSIP 终端", endpoint, "设置")

	cmd := fmt.Sprintf(QianDBA(SQL.E[101])+"|sed -n '2p' ", endpoint)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	out = out[:len(out)-1]

	if string(out) == "1" {
		fmt.Printf("record of %s have existed!\n", endpoint)
		return
	}

	cmd = fmt.Sprintf(QianDBA(SQL.E[1]), endpoint)
	out, _ = exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]
	fmt.Println(string(out))

	// username 和 终端id 保持一致
	username := endpoint
	cmd = fmt.Sprintf(QianDBA(SQL.E[2]), endpoint, "userpass", "hardly to be trusted", username)
	out, _ = exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]
	fmt.Println(string(out))

	aors, auth := endpoint, endpoint
	cmd = fmt.Sprintf(QianDBA(SQL.E[3]), endpoint, "transport-udp", aors, auth, "sets", "all", "ulaw", "no")
	out, _ = exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]
	fmt.Println(string(out))
}
