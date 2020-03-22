package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/palmergs/tokensearch"
	"os"
)

var root = tokensearch.NewTokenNode()

func main() {
	importFile := flag.String("f", "", "prepopulate with JSON file")
	histoCount := flag.Int("h", 0, "number of histogram items to return (instead of JSON)")
	flag.Parse()

	fmt.Printf("Prepopulate tree with %s...", *importFile)
	root.InsertFromFile(*importFile)
	fmt.Printf("...done\n")

	reader := bufio.NewReader(os.Stdin)
	pool := tokensearch.NewTokenNodeVisitorPool(root)
	pool.AdvanceThrough(reader)

	if *histoCount > 0 {
		histo := NewHistoEntries(pool.Matches)
		histo.PrintList(*histoCount, true)
	} else {
		writer := bufio.NewWriter(os.Stdout)
		defer writer.Flush()
		json.NewEncoder(writer).Encode(pool.Matches)
	}
}
