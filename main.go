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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	a, err := decodeMsg(string(bodyBytes))
	if err != nil {
		log.Println("decodeMsg err", err)
		return
	}
	fmt.Printf("%#v\n", a)

	s := fmt.Sprintf("%v", a)

	var reply string
	if *receiver != "" {
		// enable this for testing
		reply, err = SendPerson(s, "wenzhenglin")
	} else {
		reply, err = Send(s)
	}
	if err != nil {
		log.Println("SendPerson err", err)
		return
	}
	log.Println("done: ", reply)

	fmt.Fprintf(w, "got it")
}

// func pretty(prefix string, a interface{}) {
// 	b, _ := json.MarshalIndent(a, "", "  ")
// 	fmt.Println(prefix, string(b))
// }
