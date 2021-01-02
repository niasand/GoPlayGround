package GoExamplesReadOnly

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	var ipaddr string
	ipaddrLen := len(ip)
	for i := 0; i < ipaddrLen; i++ {
		ipaddr += fmt.Sprintf("%v.", ip[i])

	}
	ipaddr = ipaddr[:len(ipaddr)-1]
	return ipaddr
}
