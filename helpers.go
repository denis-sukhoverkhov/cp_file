package cpfile

import (
	"log"
	"os"
)

func createTestFile(name string, sizeInBytes int) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if sizeInBytes != 0 {
		err = Lorem(sizeInBytes, file)
		if err != nil {
			log.Fatal(err)
		}
	}
}
