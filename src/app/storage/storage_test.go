package storage_test

import (
	"app/storage"
	"log"
	"testing"
)

func TestExecutionQueries(t *testing.T) {
	storage.Init()

	results, err := storage.GetDagExecutionsByName()

	if err != nil {
		t.Error(err)
	}

	for _, result := range results {
		log.Println(result)
	}
}
