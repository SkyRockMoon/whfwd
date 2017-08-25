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
	webhookData := make(map[string]interface{})
	//pullRequestData := make(map[string]interface{})
	json.NewDecoder(r.Body).Decode(&webhookData)
	//json.NewDecoder(r.Body).Decode(&pullRequestData)

	fmt.Println("got webhook payload: ")
	for key, value := range webhookData {
		fmt.Printf("%s : %v\n", key, value)
	}

	fmt.Println("=========================================================\n")
	bodyBytes,_ := ioutil.ReadAll(r.Body)
    	//bodyString := string(bodyBytes)	
	fmt.Println(bodyBytes)

	//var f interface{}
	//json.Unmarshal(r.Body, &f)
	//m := f.(map[string]interface{})
	//for key, value := range m {
	//	fmt.Printf("%s : %v\n",key, value)
	//}
	//fmt.Println(pullRequestData)
}

func main() {
    log.Println("server started at http://localhost:8080")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
