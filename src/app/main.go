package main

import (
	"app/common"
	"app/orchestration"
	"app/storage"
	"app/webserver"

	"time"
)

func main() {
	common.InitialPrint()
	common.SymphoniesFolder = "../../" + common.SymphoniesFolder // set by webserver if not found

	go storage.Init()
	executionLogger := orchestration.BeginConcert()

	go storage.LogExecutions(executionLogger)
	go webserver.StartServer()

	go forever()
	select {} // block forever
}

func forever() {
	for {
		common.LogMemoryUsage()
		time.Sleep(time.Second * 60)
	}
}
