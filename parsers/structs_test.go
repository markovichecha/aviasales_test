package parsers

var expectedResult Hotel = Hotel{
	Name:        "Little Italy",
	Description: "Looking Good!",
	Address:     "Pushkin str. 1",
	City:        "Moscow",
	Country:     "Russia",
	CountryCode: "RU",
	Longitude:   float64(12.3456789),
	Latitude:    float64(98.7654321),
	StarRating:  uint8(5),
	Images:      []string{"url_1", "url_2", "url_3", "url_4", "url_5"},
	Rating:      float64(4.5),
}
