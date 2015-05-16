package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"encoding/json"
	"github.com/palmergs/tokensearch"
)

var root = tokensearch.NewTokenNode()

func main() {
	importFile := flag.String("f", "", "prepopulate with file")
	histoCount := flag.Int("h", 0, "number of histogram items to return (instead of JSON)")
	flag.Parse()

	fmt.Printf("Prepopulate tree with %s...\n", *importFile)
	root.InsertFromFile(*importFile)

	reader := bufio.NewReader(os.Stdin)
	pool := tokensearch.NewTokenNodeVisitorPool(root)
	pool.AdvanceThrough(reader)

	if *histoCount > 0 {
		histo := NewHistoEntries(pool.Matches)
		histo.PrintList(*histoCount)
	} else {
		writer := bufio.NewWriter(os.Stdout)
		defer writer.Flush()
		json.NewEncoder(writer).Encode(pool.Matches)
	}
}