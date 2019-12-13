package main

import (
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/markovichecha/aviasales_test/config"
	"github.com/markovichecha/aviasales_test/parsers"
	"github.com/markovichecha/aviasales_test/storage"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	configFile  string
	workingDir  string
	workingFile string
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

func storeData(dbs *storage.DBstorage, hotelChan chan parsers.Hotel, stop chan bool) {
	var wg sync.WaitGroup
	for {
		if hotel, ok := <-hotelChan; ok {
			wg.Add(1)
			go dbs.StoreHotel(hotel, &wg)
		} else {
			break
		}
	}
	wg.Wait()
	stop <- true
}

func init() {
	flag.StringVar(&configFile, "config", "config.yml", "path to your .yml config")
	flag.StringVar(&workingDir, "dir", "", "path to a dumps folder")
	flag.StringVar(&configFile, "file", "", "path to your dump")
	flag.StringVar(&configFile, "c", "config.yml", "path to your .yml config")
	flag.StringVar(&workingDir, "d", "", "path to a dumps folder")
	flag.StringVar(&configFile, "f", "", "path to your dump")
}

func main() {
	flag.Parse()
	var wg sync.WaitGroup
	hotelChan := make(chan parsers.Hotel)
	cfg := config.NewConfig(configFile)
	dbs := storage.NewDbStorage(cfg)
	stop := make(chan bool)
	go storeData(dbs, hotelChan, stop)
	if workingDir != "" {
		files, err := ioutil.ReadDir(workingDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			wg.Add(1)
			path := fmt.Sprintf("%s/%s", workingDir, file.Name())
			go parseDump(path, &wg, hotelChan)
		}
	} else if workingFile != "" {
		wg.Add(1)
		go parseDump(workingFile, &wg, hotelChan)
	} else {
		log.Fatal("There's nothing to dump")
	}
	wg.Wait()
	close(hotelChan)
	<-stop
}
