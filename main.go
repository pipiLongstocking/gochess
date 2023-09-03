package main

import (
	"fmt"
	"gochess/defs"
)

func main() {
	defs.ChessMetaInit()
	defs.PrintSq120Board()
	fmt.Printf("\n\n\n\n")
	defs.PrintSq64Board()
}
