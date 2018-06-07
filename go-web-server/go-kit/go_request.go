package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://10.202.63.9:8080/dds_schedule/v1/branch/755?from=algo&routeType=JTJB&businessType=b01"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
