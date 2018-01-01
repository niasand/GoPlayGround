package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

//URL to learn gorequest lib
var URL = "https://bitconnect.co/api/info/BTC_BCC"

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

func getR() string {
	resp, _ := http.Get(URL)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return string(body)
}

func postForm() {
	resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body), err)
}

func withGorequest() {
	request := gorequest.New()
	_, body, err := request.Get(URL).End()
	if err != nil {
		panic(err)
	}
	s, err2 := getBitCoinResponse([]byte(body))
	if err2 != nil {
		panic(err2.Error())
	}
	fmt.Println(s.Status)
	fmt.Println(s.Message)
	fmt.Println(s.Markets[0].Ask)
	fmt.Println(s.Markets[0].MarketName)

}
func withOutGoRequestPost() {
	m := map[string]interface{}{
		"name":    "jacky",
		"species": "dog",
	}
	mJson, _ := json.Marshal(m)
	contentReader := bytes.NewReader(mJson)
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

func main() {
	// getR()
	// postForm()
	withGorequest()
	withGoRequestPost()
	withOutGoRequestPost()

}
