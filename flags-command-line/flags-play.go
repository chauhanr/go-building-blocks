package main

import (
	"flag"
	"os"
)

var nl = flag.Bool("n", false, "print new line")

const(
	SPACE = " "
	NEW_LINE = "\n"
)


func main(){
	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i:=0; i< flag.NArg(); i++{
		s += flag.Arg(i)

		if *nl{
			s += NEW_LINE
		}
	}
	os.Stdout.WriteString(s)
}
