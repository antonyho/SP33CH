package main

import "flag"

var (
	dbHost = flag.String("db", "localhost", "Database connection string")
	// TODO - database credential
)

func main() {
	flag.Parse()
}
