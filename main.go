package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/markovichecha/aviasales_test/parsers"
)

func parseDump(fullPath string, wg *sync.WaitGroup, hotelChan chan parsers.Hotel) {
	file, err := os.Open(fullPath)
	if err != nil {
		log.Fatal(err)
	}
	switch filepath.Ext(file.Name()) {
	case ".csv":
		parsers.CSVParse(file, hotelChan)
	case ".json":
		parsers.JSONParse(file, hotelChan)
	case ".xml":
		parsers.XMLParse(file, hotelChan)
	}
	defer file.Close()
	defer wg.Done()
}

func storeData(hotelChan chan parsers.Hotel, stop chan bool) {
	var hotels []parsers.Hotel
	for {
		if hotel, ok := <-hotelChan; ok {
			hotels = append(hotels, hotel)
		} else {
			break
		}
	}
	fmt.Printf("%v", hotels)
	stop <- true
}

func main() {
	workingDir := "./dumps/"
	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		log.Fatal(err)
	}
	var wg *sync.WaitGroup
	var hotelChan chan parsers.Hotel
	var stop chan bool
	for _, file := range files {
		wg.Add(1)
		path := workingDir + file.Name()
		go parseDump(path, wg, hotelChan)
	}
	wg.Wait()
	close(hotelChan)
	<-stop
}
