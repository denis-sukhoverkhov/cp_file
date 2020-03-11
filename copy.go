package cpfile

import (
	"errors"
	"io"
	"log"
	"os"
)

var BUFFERSIZE = 500

func Copy(from string, to string, limit int64, offset int64) error {
	f, err := os.Open(from)
	if err != nil {
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if limit == 0 {
		fileStat, err := f.Stat()
		if err != nil {
			return err
		}
		if fileStat.Size() == 0 {
			return errors.New("fileSizeInBytes == 0")
		}
	}
	_, err = f.Seek(offset, 0)
	if err != nil {
		return err
	}

	destination, err := os.Create(to)
	if err != nil {
		return err
	}

	defer func() {
		err := destination.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	buf := make([]byte, BUFFERSIZE)
	ownLimit := int64(0)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if ownLimit+int64(n) > limit && limit > 0 {
			n = int(limit - ownLimit)
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}

		ownLimit += int64(n)
		if ownLimit >= limit {
			break
		}
	}

	return nil
}
