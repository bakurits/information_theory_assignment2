package fileio

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// SimpleRead Converts file to binary
func SimpleRead(srcFileName string, dstFileName string) {
	data, err := ioutil.ReadFile(srcFileName)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return
	}

	file, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for _, value := range data {
		newByte := fmt.Sprintf("%08s", strconv.FormatInt(int64(value), 2))
		_, err := file.WriteString(newByte)
		if err != nil {
			log.Fatalf("error occurred during writting %s", err)
			return
		}
	}
}
