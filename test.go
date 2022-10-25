package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "hello Goland!")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Printf("http server failed,err:%V\n", err)
		return
	}

}
