// Web Socket Forwarding From Bitbucket to Glip
//
// Glip does not easily integrate with bitbucket. This is a hack to make new pull request notifications avalable to glip as a webhook.
//
// Ryan Lowe
// rdlowe@ara.com
//
//
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
// Prior to runing start ngrok: ngrok http 8080
// This forwards bitbucket posts to localhost for testing. In general this is a security risk.
// open ngrok viewer: http://127.0.0.1:4040/inspect/http
//
// Configure bitbucket to post to: http://localhost:8080/webhook (via ngrok), and enamble notification on the creation of a pull request
//
// Configure glip to listen for a post in the channel you want these messages shown in 
//
// In the comments below the "incoming" request is from bitbucket and the "outgoing" request goes to glip

package main

import (
	"bytes"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"net/http"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {

	// read body of incoming post request
	data, _ := ioutil.ReadAll(r.Body)

	// grab (pretty) display name of the user making the PR	
	val,_,_,err:=jsonparser.Get(data, "pullrequest", "author", "display_name")
	display_name := string(val)

	// grab the title of the PR
	val,_,_,err=jsonparser.Get(data, "pullrequest", "title")
	title := string(val)

	// grab the link to the PR
	val,_,_,err =jsonparser.Get(data, "pullrequest", "links", "self","href")
	prUrl := string(val)
	
	// make JSON the hard way
	//body is the only required field for a glip message
	// [link](url) is a glip specific markup for an anchor
	// other fields not used include: icon and title
	message := fmt.Sprintf(`{"activity":"New Pull Request","body":"Name: %s\nTitle: %s\nLink: [Pull request](%s)\n"}`,display_name,title,prUrl)	

	// url to send outgoing post request	
	url := "https://hooks.glip.com/webhook/47c23015-f6af-452f-95bf-4052574a06da"

	// build outgoing request
	var jsonStr = []byte(message)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json") // this is a glip requirement

	// make outgoing request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// for debugging
	//fmt.Println(message)	
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	//fmt.Println("=========================================================\n")

}

func main() {
    log.Println("server started at http://localhost:8080")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
