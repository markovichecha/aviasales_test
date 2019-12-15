package parsers

// Hotel base structure, ready to be stored
type Hotel struct {
	Name        string   `xml:"name" csv:"hotel_translated_name" db:"name"`
	Description string   `xml:"descriptions>en" csv:"overview" db:"description"`
	Address     string   `xml:"address" csv:"addressline1" db:"address"`
	City        string   `xml:"city>en" csv:"city" db:"city"`
	Country     string   `xml:"country>en" csv:"country" db:"country"`
	CountryCode string   `xml:"countrytwocharcode" csv:"countryisocode" db:"country_iso"`
	Longitude   float64  `xml:"longitude" csv:"longitude" db:"longitude"`
	Latitude    float64  `xml:"latitude" csv:"latitude" db:"latitude"`
	StarRating  uint8    `xml:"stars" csv:"star_rating" db:"star_rating"`
	Images      []string `xml:"photos>photo>url" csv:"photo{1,5}"`
	Rating      float64  `xml:"-" csv:"rating_average" db:"rating"`
}
