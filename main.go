package main

import (
	"fmt"
	"github.com/Ungigdu/EventNotify/eth"
)

func main()  {
	m := make(chan string)
	eth.WatchNotification(m)
	for {
		select {
		case s:=<-m:
			fmt.Println(s)
		}
	}
}
