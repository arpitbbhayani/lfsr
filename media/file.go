package media

import (
	"io/ioutil"
	"os"
)

func ReadBytes(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

func WriteBytes(filePath string, b []byte) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		panic(err)
	}
}
