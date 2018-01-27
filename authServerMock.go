package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type authServerHandler struct {
	fileName string
	data     string
}

func main() {
	port := 8081
	//CAS
	cas := authServerHandler{fileName: "caslist"}
	fillData(&cas)
	http.HandleFunc("/GetCASList", cas.authServerHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func fillData(ash *authServerHandler) {
	pwd, _ := os.Getwd()
	txt, _ := ioutil.ReadFile(pwd + "/files/" + ash.fileName)
	ash.data = string(txt)
}

func (ash *authServerHandler) authServerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ash.data)
}
