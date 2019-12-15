package parsers

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
)

type jsonHotelImages struct {
	URL string `json:"url"`
}

type jsonHotelRating struct {
	Total float64 `json:"total"`
}

type jsonHotelInner struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Description string `json:"description"`
}

type jsonHotel struct {
	JSONHotelInner  jsonHotelInner    `json:"en"`
	CountryCode     string            `json:"country_code"`
	Longitude       float64           `json:"longitude"`
	Latitude        float64           `json:"latitude"`
	StarRating      uint8             `json:"star_rating"`
	JSONHotelImages []jsonHotelImages `json:"images"`
	JSONHotelRating jsonHotelRating   `json:"rating"`
}

func (jh jsonHotel) getImages() (images []string) {
	for _, v := range jh.JSONHotelImages {
		images = append(images, v.URL)
	}
	return
}

// ToHotel converts and validates JSONHotel into Hotel structure
func (jh jsonHotel) ToHotel() Hotel {
	return Hotel{
		Name:        jh.JSONHotelInner.Name,
		Address:     jh.JSONHotelInner.Address,
		City:        jh.JSONHotelInner.City,
		Country:     jh.JSONHotelInner.Country,
		Description: jh.JSONHotelInner.Description,
		CountryCode: jh.CountryCode,
		Longitude:   jh.Longitude,
		Latitude:    jh.Latitude,
		StarRating:  jh.StarRating,
		Images:      jh.getImages(),
		Rating:      jh.JSONHotelRating.Total,
	}
}

// JSONParse converts the JSON dump into a Hotel structure
func JSONParse(file io.Reader, hotels chan Hotel) {
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
