package main

import (
	"flag"
	"bufio"
	"os"
	"fmt"
	"io"
)

var numbered = flag.Bool("n", false, "print line number")
var slice = flag.Bool("s", false, "print using slice")

func main(){
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin), false)
	}

	for i := 0; i < flag.NArg(); i++{
		f, err := os.Open(flag.Arg(i))
		if err != nil{
			fmt.Fprintf(os.Stderr, "%s: error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		defer f.Close()
		if *slice {
			cat2(f)
		}else{
			cat(bufio.NewReader(f), *numbered)
		}
	}

}

// cat function will read all the content of the file till the end of the file
// or if it is a file then till the end of the file.
// if numbered flag is set as true then line numbering will be showed otherwise plain text will be shown.
func cat(r *bufio.Reader, numbered bool){

	for i :=1; ;i++ {
		buf, err := r.ReadBytes('\n')
		if err != nil && err != io.EOF{
			break
		}
		if numbered{
			fmt.Fprintf(os.Stdout, "%d. %s",i, buf)
		}else{
			fmt.Fprintf(os.Stdout, "%s",buf)
		}

		if err == io.EOF{
			break
		}
	}
	return
}

// cat2 function will take a file pointer as an input and read the the file
// 512 bytes at a time and print the buffer to the os.stdout. this is a way of reading a file
// byte by byte. This is a good way for dealing with files that are large and not spend too much on memory.
func cat2(f *os.File){

	const NBUF = 512
	var buf [NBUF]byte

	for i := 1;; i++{
		switch nr, err := f.Read(buf[:]); true{
		case nr <0 :
			fmt.Fprintf(os.Stderr, "cat2 : error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0: // EOF do nothing return
			return
		case nr > 0:
			//fmt.Fprintf(os.Stdout, "byte %d.\n", i)
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr{
				fmt.Fprintf(os.Stderr, "cat : error writing to stdout : %s\n", ew.Error())
			}
		}
	}

}