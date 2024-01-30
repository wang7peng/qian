package consts

var E = map[int]string{
	1: "Insert into asterisk.ps_aors (id, max_contacts) values ('%s', %d);",
	2: "Insert into asterisk.ps_auths (id, auth_type, password, username)" +
		" values ('%s', '%s', '%s', '%s');",
	3: "Insert into asterisk.ps_endpoints (id, transport, aors, auth, context, disallow, allow, direct_media) " +
		" values ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",

	// 查找
	101: "SELECT count(*) FROM asterisk.ps_aors WHERE id='%s' LIMIT 100;",
	102: "SELECT count(*) FROM asterisk.ps_auths WHERE id='%s' LIMIT 100;",
	103: "SELECT count(*) FROM asterisk.ps_endpoints WHERE id='%s' LIMIT 100;",
	104: "SELECT id FROM asterisk.ps_aors;",
	105: "SELECT id,auth_type,password,username FROM asterisk.ps_auths;",
	106: "SELECT id,transport,aors,auth,context,disallow,allow FROM asterisk.ps_endpoints;",

	// 删除
	201: "DELETE FROM asterisk.ps_aors WHERE id='%s' ;",
	202: "DELETE FROM asterisk.ps_auths WHERE id='%s' ;",
	203: "DELETE FROM asterisk.ps_endpoints WHERE id='%s' ;",
}
