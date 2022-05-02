package main

// gronit
// [go]cron[monitor]

import (
	"log"
	"os"
)

const EMPTYSTR string = ""

var opts *Options


func main() {
	args := os.Args[1:]
        api_key=abcdef12345
	buildSql("oyetoketoby80@gmail.com")
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	opts = parseOptions(args)
	serverStart(opts, db)
}
