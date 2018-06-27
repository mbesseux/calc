package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/mbesseux/calc"
)

// HandleCalc web
func HandleCalc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in calc")
	if len(r.URL.Path) < 4 {
		w.Write([]byte("usage:\n<int> <op> <int>"))
		return
	}
	params := strings.Split(r.URL.Path[1:], " ")

	opResult, err := calc.Operate(params[0], params[1], params[2])
	if err != nil {
		w.Write([]byte("Error : " + err.Error()))
	} else {
		w.Write([]byte(r.URL.Path[1:] + " = " + strconv.Itoa(opResult)))
	}

}

func main() {
	http.HandleFunc("/", HandleCalc)
	http.ListenAndServe(":2222", nil)
}
