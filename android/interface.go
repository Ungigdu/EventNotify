package android

import "github.com/Ungigdu/EventNotify/eth"

type EventDelegate interface {
	Notify(string)
}

func WatchNotification(d EventDelegate) error {
	m := make(chan string)
	err := eth.WatchNotification(m)
	if err != nil{
		return err
	}
	for {
		select {
		case s:=<-m:
			d.Notify(s)
		}
	}
	return nil
}
