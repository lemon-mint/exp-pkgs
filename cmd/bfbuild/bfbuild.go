package main

import (
	"bufio"
	"bytes"
	"flag"
	"log"
	"os"

	"v8.run/go/exp/hash/hashutil/bloom"
)

var (
	FalsePositiveRate = flag.Float64("p", 0.001, "false positive rate")
	CaseSensitive     = flag.Bool("c", false, "case sensitive")
)

func main() {
	flag.Parse()
	if flag.CommandLine.Arg(0) == "" {
		println("usage: bfbuild <file>")
		return
	}

	f, err := os.Open(flag.CommandLine.Arg(0))
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)
	var lines uint64
	for {
		_, _, err := br.ReadLine()
		if err != nil {
			break
		}
		lines++
	}

	// Seek to the beginning of the file
	f.Seek(0, 0)
	br.Reset(f)

	// Allocate bloom filter
	bf := bloom.NewBloom(lines, *FalsePositiveRate)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}
		if *CaseSensitive {
			bf.Set(line)
		} else {
			bf.Set(bytes.ToLower(line))
		}
	}

	// Open output file
	f, err = os.Create(flag.CommandLine.Arg(0) + ".blf")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer f.Close()

	// Write bloom filter to file
	b := bf.Bytes()
	if _, err := f.Write(b); err != nil {
		log.Fatalln(err)
		return
	}
}
