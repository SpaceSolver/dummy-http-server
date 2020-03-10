package main

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	PortNum string = "8088"
)

func usage() {
	fmt.Println("for Android Emulator")
	fmt.Printf("   connect http://10.0.2.2:%s\n", PortNum)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	fmt.Printf("<< request details >>\n")
	fmt.Printf("  length : %v\n", r.ContentLength)
	fmt.Printf("  method : %v\n", r.Method)
	fmt.Printf("  header : \n")
	for k, v := range r.Header {
		fmt.Printf("    [%v] : [%v]\n", k, v)
	}
	fmt.Printf("  body : \n")
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()
	fmt.Printf("%s\n", body)
}

func main() {
	usage()
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":"+PortNum, nil)
}
