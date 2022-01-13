package main

import (
	//"github.com/tchtsk/algorithms-complexity-test/src/bst"
	"github.com/tchtsk/algorithms-complexity-test/src/countingSort"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		//log.Fatal(e)
	}

	test := os.Getenv("TEST")

	bst.RunInsert()
	bst.RunSearch()
	bst.RunDelete()

	countingSort.Run()

	log.Println(test)
	log.Println("finished")
}
