package eth

import (
	"context"
	"fmt"
	"github.com/Ungigdu/EventNotify/eth/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WatchNotification(m chan string) error{
	client, err := ethclient.Dial(AccessPoint)
	if err != nil{
		return err
	}
	instance, err := generated.NewEventNotification(common.HexToAddress(EventNotificationAddress),client)
	if err != nil {
		return err
	}
	logs := make(chan *generated.EventNotificationNotify)
	sub, err := instance.WatchNotify(&bind.WatchOpts{
		Start:   nil,
		Context: context.Background(),
	},logs)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case e :=<- sub.Err():
				fmt.Println("subscript error",e)
				return
			case log := <-logs:
				m <- log.Msg
			}
		}
	}()

	return nil
}