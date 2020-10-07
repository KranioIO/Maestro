package common

import (
	"fmt"
	"log"
)

var version string = "0.0.2"
var symphoniesFolder string = "symphonies/"

// InitialPrint shows initial information about the current version of Maestro
func InitialPrint() {
	fmt.Println(" __  __                 _")
	fmt.Println("|  \\/  |               | |")
	fmt.Println("| \\  / | __ _  ___  ___| |_ _ __ ___")
	fmt.Println("| |\\/| |/ _` |/ _ \\/ __| __| '__/ _ \\")
	fmt.Println("| |  | | (_| |  __/\\__ \\ |_| | | (_) |")
	fmt.Println("|_|  |_|\\__,_|\\___||___/\\__|_|  \\___/")
	fmt.Println()

	log.Printf("Maestro has started")
	log.Printf("Version: %s", version)
	log.Printf("Symphonies Folder: %s", symphoniesFolder)
}

// SymphoniesFolder contains the path to all projects orchestrated by Maestro
var SymphoniesFolder = symphoniesFolder
