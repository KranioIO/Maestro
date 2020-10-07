package orchestration_test

import (
	"app/common"
	"app/orchestration"
	"log"
	"testing"
	"time"
)

func TestConcert(t *testing.T) {
	common.SymphoniesFolder = "../../../" + common.SymphoniesFolder
	orchestration.BeginConcert()

	for _, dag := range orchestration.Dags {
		log.Println(dag)
	}

	go forever()
	select {} // block forever
}

func forever() {
	for {
		common.LogMemoryUsage()
		time.Sleep(time.Second * 60)
	}
}
