package android

import "github.com/Ungigdu/EventNotify/eth"

type EventDelegate interface {
	Notify(string)
}

func WatchNotification(d EventDelegate)  {
	m := make(chan string)
	go eth.WatchNotification(m)
	for {
		select {
		case s:=<-m:
			d.Notify(s)
		}
	}
}
