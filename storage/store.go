package storage

import (
	"fmt"

	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/markovichecha/aviasales_test/config"
	"github.com/markovichecha/aviasales_test/parsers"
)

const (
	hotelsInsert = `INSERT INTO hotels(name, city, star_rating, rating, country, country_iso, address, latitude, longitude, description)
	VALUES(:name, :city, :star_rating, :rating, :country, :country_iso, :address, :latitude, :longitude, :description) RETURNING id`
	hotelsPhotoInsert = `INSERT INTO hotels_photo(url, hotel_id) VALUES($1, $2)`
)

// DBstorage represents a database connection pool
type DBstorage struct {
	connectionData string
	db             *sqlx.DB
}

// StoreHotel is storing hotels and hotels_photo
func (pool *DBstorage) StoreHotel(hotel parsers.Hotel, wg *sync.WaitGroup) {
	pool.openConnection()
	defer pool.closeConnection()
	defer wg.Done()
	id := pool.storeHotel(hotel)
	pool.storeHotelPhotos(hotel.Images, id)
}

func (pool *DBstorage) storeHotel(hotel parsers.Hotel) int {
	rows, err := pool.db.NamedQuery(hotelsInsert, hotel)
	if err != nil {
		log.Fatal(err)
	}
	var id int
	if rows.Next() {
		rows.Scan(&id)
	}
	return id
}

func (pool *DBstorage) storeHotelPhotos(photos []string, hotelID int) {
	tx := pool.db.MustBegin()
	for _, url := range photos {
		tx.MustExec(hotelsPhotoInsert, url, hotelID)
	}
	tx.Commit()
}

// NewDbStorage reads the DbStorage config for future connections
func NewDbStorage(cfg config.Config) *DBstorage {
	result := DBstorage{}

	pgInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName)
	result.connectionData = pgInfo

	return &result
}

func (pool *DBstorage) openConnection() {
	db, err := sqlx.Open("pgx", pool.connectionData)
	if err != nil {
		log.Fatal(err)
	}
	pool.db = db
}

func (pool *DBstorage) closeConnection() {
	if pool.db == nil {
		log.Fatal("No connection")
	}
	pool.db.Close()
}
