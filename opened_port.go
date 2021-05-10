package pscan

import (
	"fmt"
	"net"
	"time"
)

func OpenedPort(address string, port int32) (opened bool) {
	con, err := net.DialTimeout("tcp",
		fmt.Sprintf("%s:%d", address, port),
		5*time.Second)
	if err == nil {
		opened = true
		con.Close()
	}
	return
}
