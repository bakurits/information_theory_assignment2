package lpz

import (
	"bufio"
	"fmt"
	"information_theory_assignment2/fileio"
	"io"
	"strconv"
)

func writeGammaCode(w io.Writer, n int64) {
	binary := strconv.FormatInt(n, 2)
	for i := 0; i < len(binary); i++ {
		_, _ = w.Write([]byte{'1'})
	}
	_, _ = w.Write([]byte{'0'})
	if len(binary) > 1 {
		_, _ = w.Write([]byte(binary[1:]))
	}
}

func Compress(r io.Reader, w io.Writer, fileSize int64) {

	writeGammaCode(w, fileSize)


	pr, pw := io.Pipe()
	go func() {
		defer pw.Close()
		fileio.CompleteRead(r, pw)
	}()

	rb := bufio.NewReader(pr)
	for {
		b, err := rb.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(b)

	}
}
