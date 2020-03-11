package main

import (
	"bytes"
	//"encoding/json"
	"fmt"
	"net/http"
)

const (
	PortNum string = "8088"
)

type Ping struct {
	Status int
	Result string
}

func usage() {
	fmt.Println("for Android Emulator")
	fmt.Printf("   connect http://10.0.2.2:%s\n", PortNum)
}

func handler(w http.ResponseWriter, r *http.Request) {
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

	// json.Marshal only can parse Public member (its name must start w/ capital charactor)

	//ping := Ping{Status: http.StatusOK, Result: "ok"}
	//res, err := json.Marshal(ping)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//fmt.Printf(">>> server response: [%s]\n", res)
	//w.Write(res)

	res := "{num:\"123\"}"
	w.Write([]byte(res))
	fmt.Printf(">>> server response: [%s]\n", res)
}

func main() {
	usage()
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":"+PortNum, nil)
}
