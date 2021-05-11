package pscan

import (
	"fmt"
	"io"
)

const (
	PortMin        = 1
	PortMax        = 65545
	ConcurrencyMax = 100
)

func Pscan(address string, concurrency int, out io.Writer) {
	if concurrency > ConcurrencyMax {
		concurrency = ConcurrencyMax
	}
	scan(PortMin, 0, concurrency, address, out)
}

func scan(last, finished, concurrency int, address string, out io.Writer) {
	if finished >= PortMax {
		return
	}
	loopStop := last + concurrency
	c := make(chan int)
	for i := last; i < loopStop; i++ {
		go func(port int, c chan int) {
			opened := OpenedPort(address, port)
			if opened {
				fmt.Fprintf(out, "%d\n", port)
			}
			c <- port
		}(i, c)
	}
	for i := last; i < loopStop; i++ {
		<-c
		finished++
	}
	close(c)
	scan(loopStop, finished, concurrency, address, out)
}
