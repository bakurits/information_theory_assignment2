package fileio

import (
	"bufio"
	"io"
)

func readBytes(rb *bufio.Reader) []byte {
	var res []byte
	for ; len(res) < 8; {
		value, err := rb.ReadByte()
		if err != nil {
			break
		}
		res = append(res, value)
	}
	return res
}

// CompleteWrite converts binary to file
func CompleteWrite(r io.Reader, w io.Writer) {
	rb := bufio.NewReader(r)
	wb := bufio.NewWriter(w)
	var dataLen = 0
	for ; ; {
		dataLen++
		data := readBytes(rb)
		if len(data) < 8 {
			var leftBytes = 8 - (len(data) % 8)
			data = append(data, 49)
			for i := 1; i < leftBytes; i++ {
				data = append(data, 48)
			}
			newByte, err := ByteArrayToByte(data)
			var arr = []byte{newByte}
			if err != nil {
				return
			}
			_, _ = wb.Write(arr)
			break
		}

		newByte, err := ByteArrayToByte(data)
		var arr = []byte{newByte}
		if err != nil {
			return
		}
		_, _ = wb.Write(arr)

	}
}
