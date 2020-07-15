package main

import (
	"fmt"
	"github.com/Ungigdu/EventNofify/eth"
)

func main()  {
	m := make(chan string)
	go eth.WatchNotification(m)
	for {
		select {
		case s:=<-m:
			fmt.Println(s)
		}
	}
}
