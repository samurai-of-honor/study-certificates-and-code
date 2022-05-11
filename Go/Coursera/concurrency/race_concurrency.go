package main

import (
	"fmt"
	"time"
)

func Inc(i *int) {
	*i++
}

func Dec(i *int) {
	*i--
}

func main() {
	i := 0
	tick := time.NewTicker(time.Second)

	for {
		select {
		case <-tick.C:
			go Inc(&i)
			go Dec(&i)
			fmt.Println(i)
		}
	}
}

/*
	Race condition is an undesirable situation that occurs when a device
	or system attempts to perform two or more operations at the same time,
	but because of the nature of the device or system, the operations
	must be done in the proper sequence to be done correctly.

	Race condition occur when two computer program processes,
	or threads, attempt to access the same resource at the same time
	and cause problems in the system.
*/
