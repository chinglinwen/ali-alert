package main

import (
	"fmt"
	"testing"
)

/*
alertName=test-独立报警状态码&alertState=OK&curValue=0&dimensions={taskId=b47607e8-9826-4ba0-8f04-d27c8ca0670b, userId=1627690583007430}&expression=$ErrorCodeMaximum>=201&instanceName=taskName=test，address=http://a.clwen.com:9000&metricName=StatusCode&metricProject=acs_networkmonitor&namespace=acs_networkmonitor&preTriggerLevel=INFO&ruleId=b47607e8-9826-4ba0-8f04-d27c8ca0670b_StatusCode&signature=dVHyXQ6g7jwbeHFSD/g4BcmU/Dg=&timestamp=1560334740000&triggerLevel=OK&userId=1627690583007430



*/

var a = `
	alertName=test-%E7%8B%AC%E7%AB%8B%E6%8A%A5%E8%AD%A6%E7%8A%B6%E6%80%81%E7%A0%81&alertState=OK&curValue=0&dimensions=%7BtaskId%3Db47607e8-9826-4ba0-8f04-d27c8ca0670b%2C+userId%3D1627690583007430%7D&expression=%24ErrorCodeMaximum%3E%3D201&instanceName=taskName%3Dtest%EF%BC%8Caddress%3Dhttp%3A%2F%2Fa.clwen.com%3A9000&metricName=StatusCode&metricProject=acs_networkmonitor&namespace=acs_networkmonitor&preTriggerLevel=INFO&ruleId=b47607e8-9826-4ba0-8f04-d27c8ca0670b_StatusCode&signature=dVHyXQ6g7jwbeHFSD%2Fg4BcmU%2FDg%3D&timestamp=1560334740000&triggerLevel=OK&userId=1627690583007430`

func TestDecodeMsg(t *testing.T) {
	// fmt.Println(url.QueryUnescape(a))
	// spew.Dump(url.ParseQuery(a))
	v, err := decodeMsg(a)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Printf("%#v\n", v)
}
func TestFormat(t *testing.T) {
	v, _ := decodeMsg(a)
	s := fmt.Sprintf("%v", v)
	fmt.Printf("%v\n", s)
}

func TestParseInstanceName(t *testing.T) {
	a := "taskName=test，address=http://a.clwen.com:9000"
	task, addr := parseInstanceName(a)
	if task != "test" {
		t.Error("expect task: test, got", task)
	}

	if addr != "http://a.clwen.com:9000" {
		t.Error("expect addr: http://a.clwen.com:9000, got", addr)
	}
}
func TestParseTimestamp(t *testing.T) {
	a := "1560334740000"
	// a := "1560334740"
	tm, err := parseTimestamp(a)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tm)
	// fmt.Println(tm.Local())
}
