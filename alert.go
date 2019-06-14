package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"text/template"
	"time"
)

/*
"instanceName":["taskName=test，address=http://a.clwen.com:9000"]
"preTriggerLevel":["INFO"]
"signature":["dVHyXQ6g7jwbeHFSD/g4BcmU/Dg="]
"\n\talertName":["test-独立报警状态码"]
"alertState":["OK"]
"metricProject":["acs_networkmonitor"]
"timestamp":["1560334740000"]
"metricName":["StatusCode"]
"namespace":["acs_networkmonitor"]
"triggerLevel":["OK"]
"userId":["1627690583007430"]
"curValue":["0"]
"dimensions":["{taskId=b47607e8-9826-4ba0-8f04-d27c8ca0670b, userId=1627690583007430}"]
"expression":["$ErrorCodeMaximum>=201"]
"ruleId":["b47607e8-9826-4ba0-8f04-d27c8ca0670b_StatusCode"]
*/

type Alert struct {
	// InstanceName    string
	AlertName       string
	TaskName        string
	Addr            string
	CurValue        string
	Expr            string
	AlertState      string //metricName
	PreTriggerLevel string
	Timestamp       string
	MetricName      string
	TriggerLevel    string
	// userid
}

func decodeMsg(body string) (a *Alert, err error) {
	if body == "" {
		err = fmt.Errorf("empty body")
		return
	}
	// fmt.Println("body", body)
	// sbody, err := url.QueryUnescape(body)
	// if err != nil {
	// 	return
	// }
	// fmt.Println("sbody", sbody)
	data, err := url.ParseQuery(body)
	if err != nil {
		err = fmt.Errorf("parse uri err: %v, body: %v", err, body)
		return
	}

	s := make(map[string]string)
	for i, v := range data {
		if len(v) >= 1 {
			// var k1, v1 string
			if i == "\n\talertName" {
				i = "alertName"
			}
			if i == "expression" {
				v[0] = strings.TrimPrefix(v[0], "$")
			}
			s[i] = v[0]
			// fmt.Printf("%q:%q\n", i, v)
		}
	}
	task, addr := parseInstanceName(s["instanceName"])
	tm, err := parseTimestamp(s["timestamp"])
	if err != nil {
		return
	}
	a = &Alert{
		AlertName:       s["alertName"],
		PreTriggerLevel: s["preTriggerLevel"],
		AlertState:      s["alertState"],
		Timestamp:       tm,
		MetricName:      s["metricName"],
		TriggerLevel:    s["triggerLevel"],
		CurValue:        s["curValue"],
		Expr:            s["expression"],
		TaskName:        task,
		Addr:            addr,
	}

	return
}

// "taskName=test，address=http://a.clwen.com:9000"
func parseInstanceName(i string) (task, addr string) {
	s := strings.Split(i, "，")
	if len(s) != 2 {
		return
	}
	task = strings.TrimPrefix(s[0], "taskName=")
	addr = strings.TrimPrefix(s[1], "address=")
	return
}

const layout = "2006-1-2 15:04:05"

func parseTimestamp(s string) (t string, err error) {
	ss := strings.TrimSuffix(s, "000")
	i, err := strconv.ParseInt(ss, 10, 64)
	if err != nil {
		return
	}
	tm := time.Unix(i, 0)
	t = tm.Local().Format(layout)
	return
}

/*

&main.Alert{AlertName:"test-???????", TaskName:"test", Addr:"http://a.clwen.com:9000", CurValue:"0", Expr:"ErrorCodeMaximum>=201", AlertState:"OK", PreTriggerLevel:"INFO", Timestamp:"2019-6-14 14:00:00", MetricName:"StatusCode", TriggerLevel:"OK"}

*/

const msgtmpl = `报警名字： {{ .AlertName }}
任务： {{ .TaskName }}
地址： {{ .Addr }}
当前值： {{ .CurValue }}
条件： {{ .MetricName }}: {{ .Expr }}
报警状态： {{ .AlertState }}
时间： {{ .Timestamp }}`

func (a *Alert) String() (s string) {
	t, err := template.New("test").Parse(msgtmpl)
	if err != nil {
		log.Fatal("template error", err)
	}
	var b bytes.Buffer
	err = t.Execute(&b, a)
	s = b.String()
	return
}
