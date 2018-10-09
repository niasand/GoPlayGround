package main

import (
	"github.com/gin-gonic/gin/json"
	"os"
)

type Server struct {
	ID int `json:"-"`

	ServerName  string `json:"servername"`
	ServerName2 string `json:"serverName2,string"`

	ServerIP string `json:"serverIP,omitempty"`
}

func main() {

	s := Server{
		ID:          3,
		ServerName:  `Go "1.0"`,
		ServerName2: `Go "1.0"`,
		ServerIP:    ``,
	}

	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}
