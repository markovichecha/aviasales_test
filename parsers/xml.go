package parsers

import (
	"encoding/xml"

	"io"
	"log"
	"os"
)

// XMLParse converts the XML dump into a Hotel structure
func XMLParse(file *os.File, hotels chan Hotel) {
	dec := xml.NewDecoder(file)

	for {
		t, err := dec.Token()

		if t == nil || err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("Error decoding token: %s", err)
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "hotel":
				var hotel Hotel

				if err = dec.DecodeElement(&hotel, &se); err != nil {
					log.Fatalf("Error decoding item: %s", err)
				}

				hotels <- hotel
			}
		}
	}
}
