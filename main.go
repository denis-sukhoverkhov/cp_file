package main

import (
	cpfile "cpfile/src"
	"flag"
	"log"
)

func main() {

	var pathToCopyFile = flag.String("pathToCopyFile", "",
		"path to file which need to copy")
	var pathToSaveFile = flag.String("pathToSaveFile", "",
		"path for saving file")
	var limit = flag.Int64("limit", 0, "max bytes which need copied")
	var offset = flag.Int64("offset", 0, "offset in bytes from head of file")

	flag.Parse()

	err := cpfile.Copy(*pathToCopyFile, *pathToSaveFile, *limit, *offset)
	if err != nil {
		log.Fatal(err)
	}
}
