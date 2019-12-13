package parsers

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

// JSONParse converts the JSON dump into a Hotel structure
func JSONParse(file *os.File, hotels chan Hotel) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var jh jsonHotel
		err := json.Unmarshal(scanner.Bytes(), &jh)
		if err != nil {
			log.Fatal(err)
		}
		hotels <- jh.ToHotel()
	}
}
