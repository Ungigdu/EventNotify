SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	solc EventNotification.sol --abi --overwrite -o ./
	abigen --abi EventNotification.abi --pkg generated --type EventNotification --out EventNotification.go
