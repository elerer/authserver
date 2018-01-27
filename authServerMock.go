package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type casListHandler struct {
	data string
}

func main() {
	port := 8081
	cas := newCasListHandler()
	http.HandleFunc("/GetCASList", cas.serveCasHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func newCasListHandler() *casListHandler {
	cas := new(casListHandler)
	pwd, _ := os.Getwd()
	txt, _ := ioutil.ReadFile(pwd + "/files/casList")
	cas.data = string(txt)
	return cas
}

func (cas *casListHandler) serveCasHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, cas.data)
}
