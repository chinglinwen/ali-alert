package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	flag.Parse()
	log.Println("started, ready to receive callback call")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("empty body err: %v", err)
		E(w, err)
		return
	}

	a, err := decodeMsg(string(bodyBytes))
	if err != nil {
		err = fmt.Errorf("decodeMsg err: %v", err)
		E(w, err)
		return
	}
	fmt.Printf("got callback: %v\n", a)

	s := fmt.Sprintf("%v", a)

	var reply string
	if *receiver != "" {
		// enable this for testing
		reply, err = SendPerson(s, "wenzhenglin")
	} else {
		reply, err = Send(s)
	}
	if err != nil {
		err = fmt.Errorf("send err: %v, reply: %v", err, reply)
		E(w, err)
		return
	}
	log.Println("send reply: ", reply)

	fmt.Fprintf(w, "got it")
}

func E(w http.ResponseWriter, err error) {
	log.Println(err)
	fmt.Fprintf(w, err.Error())
}

// func pretty(prefix string, a interface{}) {
// 	b, _ := json.MarshalIndent(a, "", "  ")
// 	fmt.Println(prefix, string(b))
// }
