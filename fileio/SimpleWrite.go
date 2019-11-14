package fileio

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// SimpleWrite converts binary to file
// works only for files with length of multiples 8
func SimpleWrite(srcFileName string, dstFileName string) {
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
		err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	for i := 0; i < len(data); i += 8 {
		newByte, err := ByteArrayToByte(data[i : i+8])
		var arr = []byte{newByte}
		if err != nil {
			return
		}
		_, _ = file.Write(arr)
	}
}
