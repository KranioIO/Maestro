package storage

import (
	"app/common"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// GetAllSymphoniesFilenamesInFilesystem returns a list of filesystem paths with the names of all symphonies
func GetAllSymphoniesFilenamesInFilesystem() []string {
	dir := common.SymphoniesFolder
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	var pipelinesFiles []string

	for _, file := range files {
		if name := file.Name(); strings.Contains(name, ".yml") {
			pipelinesFiles = append(pipelinesFiles, dir+name)
		}

	}

	return pipelinesFiles
}

// ReadYmlFile parse the local yml file to Dict type
func ReadYmlFile(filePath string) common.Dict {
	ymlFile, _ := os.Open(filePath)
	ymlByteArray, _ := ioutil.ReadAll(ymlFile)
	ymlMap := make(common.Dict)

	err := yaml.Unmarshal(ymlByteArray, &ymlMap)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return ymlMap
}
