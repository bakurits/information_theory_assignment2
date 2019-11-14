package fileio

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

// SimpleRead Converts file to binary
func SimpleRead(r io.Reader, w io.Writer) {

	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
		return
	}
	for _, value := range data {
		_, err := fmt.Fprintf(w, "%08s", strconv.FormatInt(int64(value), 2))
		if err != nil {
			log.Fatalf("error occurred during writting %s", err)
			return
		}
	}
}
