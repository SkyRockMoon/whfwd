package main

import (
	"fmt"
	//"io"
	"log"
	"net/http"
	//"os"
	"encoding/json"
	"io/ioutil"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=========================================================\n")
	bodyBytes,_ := ioutil.ReadAll(r.Body)
	//fmt.Println(bodyBytes)
    	bodyString := string(bodyBytes)	
	fmt.Println(bodyString)
	fmt.Println("=========================================================\n")

	webhookData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&webhookData)

	fmt.Println("got webhook payload: ")
	for key, value := range webhookData {
		fmt.Printf("%s : %v\n", key, value)
	}
	fmt.Println("=========================================================\n")


	var f interface{}
	json.Unmarshal(bodyBytes, &f)
	m := f.(map[string]interface{})
	for key, value := range m {
		fmt.Printf("%s : %v\n",key, value)
	}
//	fmt.Println(pullRequestData)
}

func main() {
    log.Println("server started at http://localhost:8080")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
