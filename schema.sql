CREATE TABLE IF NOT EXISTS hotels (
	id SERIAL NOT NULL PRIMARY KEY,
    name TEXT,
    city TEXT,
    star_rating INT,
    rating FLOAT,
    country TEXT,
	country_iso VARCHAR(3),
	address TEXT,
	latitude FLOAT,
	longitude FLOAT,
	description TEXT
);

CREATE TABLE IF NOT EXISTS hotels_photo (
	id SERIAL NOT NULL PRIMARY KEY,
	url text,
	hotel_id SERIAL NOT NULL REFERENCES hotels(id)
);