package lpz

import (
	"fmt"
	"information_theory_assignment2/fileio"
	"io"
	"log"
	"math"
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
	_, _ = w.Write([]byte{'\n'})
}

type trieNode struct {
	left  *trieNode
	right *trieNode
	index int
}

type Lexicon struct {
	size int
	root *trieNode
}

func (lex *Lexicon) MaxIndexSize() int {
	logLow := math.Log2(float64(lex.size))
	res := int(logLow)
	if logLow-float64(res) < 0.0000001 {
		return res
	} else {
		return res + 1
	}

}

// getMaxMatch finds biggest match of given input in lexicon
// returns biggest match node
// if input ended returns true
func getMaxMatch(r io.Reader, lexicon Lexicon) (*trieNode, bool) {
	curNode := lexicon.root
	hasEnded := false
	for {
		curByte := make([]byte, 1)

		if !hasEnded {

			_, err := r.Read(curByte)
			if err != nil {
				hasEnded = true
			}
		} else {
			curByte[0] = '0'
		}

		if curNode.index > -1 {
			break
		}

		if curByte[0] == '0' {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}
	}

	return curNode, hasEnded
}

func WriteInFile(w io.Writer, value int, outputSize int) {
	format := fmt.Sprintf("%%0%ds\n", outputSize)
	_, err := fmt.Fprintf(w, format, strconv.FormatInt(int64(value), 2))
	if err != nil {
		log.Panic(err)
	}
}

func enlargeLexicon(lexicon *Lexicon, node *trieNode) {
	node.left = &trieNode{
		left:  nil,
		right: nil,
		index: node.index,
	}
	node.right = &trieNode{
		left:  nil,
		right: nil,
		index: lexicon.size,
	}
	node.index = -1
	lexicon.size++
}

func getInitialLexicon() Lexicon {
	return Lexicon{
		size: 2,
		root: &trieNode{
			left: &trieNode{
				left:  nil,
				right: nil,
				index: 0,
			},
			right: &trieNode{
				left:  nil,
				right: nil,
				index: 1,
			},
			index: -1,
		},
	}
}

func Compress(r io.Reader, w io.Writer, fileSize int64) {

	writeGammaCode(w, fileSize)

	pr, pw := io.Pipe()
	go func() {
		fileio.SimpleRead(r, pw)
		_ = pw.Close()
	}()

	lexicon := getInitialLexicon()

	for {

		trieNode, hasEnded := getMaxMatch(pr, lexicon)

		if trieNode == nil || trieNode.index == -1 {
			panic("node isn't correct")
		}

		WriteInFile(w, trieNode.index, lexicon.MaxIndexSize())
		enlargeLexicon(&lexicon, trieNode)

		if hasEnded {
			break
		}

	}

}
