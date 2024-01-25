package guidebook

func Run(arg2 string) {
	switch arg2 {
	case "48":
		CreateEndPointA()
		CreateEndPointB()
	case "68", "69":
		// 68 新增两个话机的就是为了区分分机的 transport
		ShowAll_EndPoints()
		// 另外插入两条记录
		CreateEndPoint68("SOFTPHONE_A", "transport-tls")
		CreateEndPoint68("SOFTPHONE_B", "transport-tls")

		CreateAors69("SOFTPHONE_A", 2)
		CreateAors69("SOFTPHONE_B", 2)
		CreateAuths69("SOFTPHONE_A", "userpass", "iwouldnotifiwereyou")
		CreateAuths69("SOFTPHONE_B", "userpass", "areyoueventrying")
		ShowAll_Auths()

	default:
	}
}
