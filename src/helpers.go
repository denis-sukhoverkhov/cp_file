package cpfile

import (
	"crypto/md5"
	"encoding/hex"
	"io"
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

func GetFileSize(fileName string) int64 {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fileStat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return fileStat.Size()
}

func IsEqualFiles(f1 string, f2 string) bool {
	f, err := os.Open(f1)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	h1 := md5.New()
	if _, err := io.Copy(h1, f); err != nil {
		log.Fatal(err)
	}

	hashInBytes := h1.Sum(nil)

	hash1 := hex.EncodeToString(hashInBytes)

	ff, err := os.Open(f2)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := ff.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	h2 := md5.New()
	if _, err := io.Copy(h2, ff); err != nil {
		log.Fatal(err)
	}

	hashInBytes = h1.Sum(nil)
	hash2 := hex.EncodeToString(hashInBytes)
	return hash1 == hash2
}
