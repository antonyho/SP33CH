package main

import "flag"

var (
	dbHost = flag.String("db", "localhost", "Database connection string")
	prod   = flag.Bool("prod", true, "Flag for production instance")
	test   = flag.Bool("test", false, "Flag for testing instance")
)

func main() {
	flag.Parse()

	if *test {
		*prod = false
	}

}
