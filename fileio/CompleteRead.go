package fileio

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

// CompleteRead Converts file to binary
func CompleteRead(r io.Reader, w io.Writer) {
	rb := bufio.NewReader(r)
	wb := bufio.NewWriter(w)
	var newByte string

	for {
		value, err := rb.ReadByte()
		if err != nil {
			break
		}
		newByte = fmt.Sprintf("%08s", strconv.FormatInt(int64(value), 2))
		_, err = wb.WriteString(newByte)
		if err != nil {
			log.Fatalf("error occurred during writting %s", err)
			return
		}
	}

	var idx = strings.LastIndexByte(newByte, '1')
	if idx <= 0 {
		return
	}
	newByte = newByte[:idx]
	_, err := wb.WriteString(newByte)
	if err != nil {
		log.Fatalf("error occurred during writting %s", err)
		return
	}
}
