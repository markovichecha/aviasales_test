package parsers

// Hotel base structure, ready to be stored
type Hotel struct {
	Name        string   `xml:"name" csv:"hotel_translated_name"`
	Description string   `xml:"descriptions>en" csv:"overview"`
	Address     string   `xml:"string" csv:"addressline1"`
	City        string   `xml:"city>en" csv:"city"`
	Country     string   `xml:"country>en" csv:"country"`
	CountryCode string   `xml:"countrytwocharcode" csv:"countryisocode"`
	Longitude   float64  `xml:"longitude" csv:"longitude"`
	Latitude    float64  `xml:"latitude" csv:"latitude"`
	Stars       uint8    `xml:"stars" csv:"star_rating"`
	Images      []string `xml:"photos>photo>url" csv:"photo{1,5}"`
}

type jsonHotelImages struct {
	URL string `json:"url"`
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
	Stars           uint8             `json:"star_rating"`
	JSONHotelImages []jsonHotelImages `json:"images"`
}

func (jh jsonHotel) getImages() (images []string) {
	for _, v := range jh.JSONHotelImages {
		images = append(images, v.URL)
	}
	return
}

func (jh jsonHotel) convertStars() (stars uint8) {
	stars = uint8(jh.Stars / 10)
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
		Stars:       jh.convertStars(),
		Images:      jh.getImages(),
	}
}
