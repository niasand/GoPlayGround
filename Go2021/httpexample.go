package Go2021

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/parnurzeal/gorequest"
	"github.com/dghubble/go-twitter/twitter"
)

//URL to learn gorequest lib


//BitcoinInfo parse the response of the api
type BitcoinInfo struct {
	Bid            string `json:"bid"`
	LastPrice      string `json:"last_price"`
	Volume24h      string `json:"volume24h"`
	Currencystring string `json:"currency"`
	MarketName     string `json:"marketname"`
	Ask            string `json:"ask"`
	Low24h         string `json:"low24h"`
	Change24h      string `json:"change24h"`
	High24h        string `json:"high24h"`
	BaseCurrency   string `json:"basecurrency"`
}

//BrowserVersionSupport
type BrowserVersionSupport struct {
	Chrome  string
	Firefox string
}

//BitcoinInfoAPIResponse include the BitcoinInfo struct
type BitcoinInfoAPIResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Markets []BitcoinInfo `json:"markets"`
}

func getBitCoinResponse(body []byte) (*BitcoinInfoAPIResponse, error) {
	var s = new(BitcoinInfoAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops...", err)
	}
	return s, err
}



func postForm() {
	resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), err)
}


func withOutGoRequestPost() {
	m := map[string]interface{}{
		"name":    "jacky",
		"species": "dog",
	}
	mJSON, _ := json.Marshal(m)
	contentReader := bytes.NewReader(mJSON)
	req, _ := http.NewRequest("POST", "http://example.com", contentReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Notes", "Go request is coming! ")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.Body)
}

func withGoRequestPost() {
	request := gorequest.New()
	_, body, errs := request.Post("http://example.com").
		Set("Notes", "Go request is coming! ").Send(`{"name":"jacky","species":"dog"}`).End()
	if errs != nil {
		panic(errs)
	}
	fmt.Println("######")
	fmt.Println(body)
	fmt.Println("######")
}

func sendBrowserVersion() {
	browserVersion := BrowserVersionSupport{Chrome: "63.0", Firefox: "30.0"}
	request := gorequest.New()
	resp, body, err := request.Post("http://example.com").
		Send(browserVersion).
		Send(`Safari:"5.1.10"`).End()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp, body)
}

func mainfalse() {
	// getR()
	// postForm()
	withGoRequestPost()
	withOutGoRequestPost()
	sendBrowserVersion()

}

func TwitterMsg(){
	config := oauth1.NewConfig("0", "0")
	token := oauth1.NewToken("1", "2")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	fmt.Println(tweets, resp, err)

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("本条推来自bot", nil)
	fmt.Println(tweet, resp, err)
}
