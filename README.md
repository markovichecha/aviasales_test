# aviasales_test
## Installation
`$ git clone https://github.com/markovichecha/aviasales_test`

`$ cd aviasales_test`

`$ go build`

### PostgreSQL
Run `$ sudo docker-compose up -d` to boot up a dockerized  PostgreSQL.

## Command Line Interface
`$ ./aviasales_test -c config.yml  -d dump`

-c --config:	Path to the .yml config file

-d   --dir:    Path to the directory with xml, csv and json dumps

-f   --file:	Path to the  xml, csv or json file
