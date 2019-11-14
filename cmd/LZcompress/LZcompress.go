package main

import (
	"information_theory_assignment2/lpz"
	"io"
	"log"
	"os"
)

func main() {
	inp := "/home/bakur/go/src/information_theory_assignment2/public_tests/A/003.dat"
	outp := "/home/bakur/go/src/information_theory_assignment2/public_tests/A/003.mans"
	inpf, err := os.Open(inp)
	if err != nil {
		log.Fatal("error in opening file")
	}
	defer func() {
		err = inpf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	outpf, err := os.OpenFile(outp, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("error in opening file")
	}
	defer func() {
		err = outpf.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	info, err := inpf.Stat()
	if err != nil {
		log.Fatal("error in opening file")
	}

	lpz.Compress(io.Reader(inpf), io.Writer(outpf), info.Size())

}
