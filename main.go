package main

import (
	"github.com/tchtsk/algorithms-complexity-test/src/bst"
	"log"
	//"os"
	//"github.com/joho/godotenv"
)

func main() {
	/*
		e := godotenv.Load()
		if e != nil {
			log.Fatal(e)
		}

		test := os.Getenv("TEST")
	*/

	bst.RunInsert()
	bst.RunSearch()
	bst.RunDelete()

	log.Println()
	log.Println("finished")
}
