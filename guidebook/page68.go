package guidebook

import (
	"fmt"
	"os/exec"
	SQL "qianDev/consts"
)

func ShowAll_EndPoints() {
	cmd := QianDBA(SQL.E[106])
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	out = out[:len(out)-1]
	fmt.Println(string(out))
}

func ShowAll_Auths() {
	cmd := QianDBA(SQL.E[105])
	out, _ := exec.Command("bash", "-c", cmd).Output()
	out = out[:len(out)-1]
	fmt.Println(string(out))
}

func CreateEndPoint68(endpoint, transport string) {
	cmd := fmt.Sprintf(QianDBA(SQL.E[103])+"|sed -n '2p' ", endpoint)
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

	aors, auth := endpoint, endpoint
	cmd = fmt.Sprintf(QianDBA(SQL.E[3]),
		endpoint, transport, aors, auth, "sets", "all", "ulaw", "no")
	exec.Command("bash", "-c", cmd).Output()

	ShowAll_EndPoints()
}

func CreateAors69(endpoint string, max_contacts int) {

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

	cmd = fmt.Sprintf(QianDBA(SQL.E[1]), endpoint, max_contacts)
	exec.Command("bash", "-c", cmd).Output()

	cmd = QianDBA(SQL.E[104])
	out, err = exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	out = out[:len(out)-1]
	fmt.Println(string(out))
}

func CreateAuths69(id, auth_type, password string) {
	cmd := fmt.Sprintf(QianDBA(SQL.E[102])+"|sed -n '2p' ", id)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	out = out[:len(out)-1]

	if string(out) == "1" {
		fmt.Printf("record of %s have existed!\n", id)
		return
	}

	username := id
	cmd = fmt.Sprintf(QianDBA(SQL.E[2]), id, auth_type, password, username)
	exec.Command("bash", "-c", cmd).Output()

	ShowAll_Auths()
}
