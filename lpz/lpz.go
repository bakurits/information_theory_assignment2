package lpz

import (
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
		fileio.SimpleRead(r, pw)
		_ = pw.Close()
	}()

	for {

		var curByte byte
		_, err := fmt.Fscan(pr, &curByte)
		if err != nil {
			break
		}

	}

}
