//Prior to runing start ngrok: ngrok http 8080
// open ngrok viewer: http://127.0.0.1:4040/inspect/http

package main

import (
	"fmt"
	//"io"
	"log"
	"net/http"
	//"os"
	//"encoding/json"
	"io/ioutil"
	"github.com/buger/jsonparser"
	"bytes"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=========================================================\n")
	data,_ := ioutil.ReadAll(r.Body)
	
	val,dataType,offset,err:=jsonparser.Get(data, "pullrequest", "author", "display_name")
	display_name := string(val)
	fmt.Println(display_name, dataType, offset, err)

	val,dataType,offset,err=jsonparser.Get(data, "pullrequest", "title")
	title := string(val)
	fmt.Println(string(title), dataType, offset, err)

	val,dataType,offset,err =jsonparser.Get(data, "pullrequest", "links", "self","href")
	prUrl := string(val)
	fmt.Println(prUrl)

	fmt.Println("=========================================================\n")

	
	//message := fmt.Sprintf("{\"activity\":\"New Pull Request\",\"body\":\"Name: %s\nTitle: %s\nLink: [Pull request](%s)\n\"}",display_name,title,link)
	message := fmt.Sprintf(`{"activity":"New Pull Request","body":"Name: %s\nTitle: %s\nLink: [Pull request](%s)\n"}`,display_name,title,link)	
	fmt.Println(message)

	url := "https://hooks.glip.com/webhook/47c23015-f6af-452f-95bf-4052574a06da"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	fmt.Println("=========================================================\n")

}

func main() {
    log.Println("server started at http://localhost:8080")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
