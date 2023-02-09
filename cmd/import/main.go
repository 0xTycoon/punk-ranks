package main

import (
	"github.com/currencytycoon/punkranking"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	defer punksranking.Close()

	if _, err := punksranking.SetupDB("./config.json"); err != nil {
		log.Print(err)
		return
	}
	punksranking.Zap()
	if err := punksranking.Import("punks.csv"); err != nil {
		log.Print(err)
		return
	}
	if err := punksranking.ImportAttr(); err != nil {
		log.Print(err)
		return
	}
	if err := punksranking.Calculate(); err != nil {
		log.Print(err)
		return
	}

	if err := punksranking.Link(); err != nil {
		log.Print(err)
		return
	}

}
