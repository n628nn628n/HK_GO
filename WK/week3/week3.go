package main

import (
	"encoding/json"
	"fmt"
)

type ApiResult struct {
	Ret []struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Domain   int    `json:"domain"`
		Currency string `json:"currency"`
		Cash     struct {
			ID          int     `json:"id"`
			UserID      int     `json:"user_id"`
			Balance     float64 `json:"balance"`
			PreSub      int     `json:"pre_sub"`
			PreAdd      int     `json:"pre_add"`
			Currency    string  `json:"currency"`
			LastEntryAt string  `json:"last_entry_at"`
		} `json:"cash"`
	} `json:"ret"`
	Result  string `json:"result"`
	Profile struct {
		ExecutionTime int    `json:"execution_time"`
		ServerName    string `json:"server_name"`
	} `json:"profile"`
}

type conn struct {
	Ip   string
	Host string
	Port string
}

func main() {
	Q1()
	Q2()
}

func Q1() {
	var api_struct ApiResult
	_ = json.Unmarshal(getApiResult(), &api_struct)
	fmt.Printf("Q1: id = %v, domain = %v, username = %v, cash = %v \n",
		api_struct.Ret[0].ID, api_struct.Ret[0].Domain, api_struct.Ret[0].Username, api_struct.Ret[0].Cash)
}

func Q2() {
	/* Q2-1
	IP:192.17.55.3	Host:docs.google.com 		Port:80
	IP:192.17.33.17 Host:ts-phpadmin.vir999.com Port:80
	IP:192.17.99.52 Host:jsonviewer.stack.hu 	Port:7711
	*/
	var setList []*conn
	set1 := conn{Ip: "192.17.55.3", Host: "docs.google.com", Port: "80"}
	set2 := conn{Ip: "192.17.33.17", Host: "ts-phpadmin.vir999.com", Port: "80"}
	set3 := conn{Ip: "192.17.99.52", Host: "jsonviewer.stack.hu", Port: "7711"}

	setList = append(setList, &set1)
	setList = append(setList, &set2)
	setList = append(setList, &set3)

	fmt.Printf("Q2-1: ")
	printResult(setList)

	/* Q2-2
	IP:177.18.2.33 Host:github.com Port:80
	*/
	set4 := conn{Ip: "177.18.2.33", Host: "github.com", Port: "80"}
	setList = append(setList, &set4)

	fmt.Printf("Q2-2: ")
	printResult(setList)

	// Q2-3
	updateSetting(setList)

	fmt.Printf("Q2-3: ")
	printResult(setList)
}

func updateSetting(setList []*conn) {
	for i, set := range setList {
		if set.Host == "jsonviewer.stack.hu" {
			setList[i].Port = "80"
		}
	}
}

func getApiResult() []byte {
	userData := `{"ret":[{"id":148500286,"parent_id":1414663,"parent":1414663,"all_parents":[1414663,59735,58971,24367,6],"domain":6,"last_login":"2018-06-12T09:17:51+0800","currency":"CNY","password_expire_at":"2020-09-04T18:19:02+0800","password_reset":false,"last_bank":54,"username":"wesley01","block":false,"bankrupt":false,"cash":{"id":71814701,"user_id":148500286,"balance":4923642.0545,"pre_sub":0,"pre_add":0,"currency":"CNY","last_entry_at":"20180612155900"},"cash_fake":null,"card":null,"enabled_card":null}],"result":"ok","profile":{"execution_time":14,"server_name":"ipl-web01.rd5.qa"}}`

	return []byte(userData)
}

func printResult(result []*conn) {
	for _, set := range result {
		fmt.Printf("%+v ", *set)
	}

	fmt.Println()
}
