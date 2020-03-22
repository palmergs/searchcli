package main

import (
	"fmt"
	"github.com/palmergs/tokensearch"
	"sort"
)

type HistoEntry struct {
	ident int64
	name  string
	count int
}

type HistoEntries []*HistoEntry

func NewHistoEntries(matches []*tokensearch.TokenMatch) HistoEntries {
	histo := make(map[string]*HistoEntry)
	for _, match := range matches {
		ident := match.Token.Ident
		key := match.Token.Display
		if histo[key] == nil {
			histo[key] = &HistoEntry{ident, key, 0}
		}
		histo[key].count += 1
	}

	entries := make(HistoEntries, 0)
	for _, value := range histo {
		entries = append(entries, value)
	}
	return entries
}

func (slice HistoEntries) Len() int {
	return len(slice)
}

func (slice HistoEntries) Less(i, j int) bool {
	return slice[i].count > slice[j].count
}

func (slice HistoEntries) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice HistoEntries) PrintList(limit int, header bool) {
	sort.Sort(slice)
	if limit > len(slice) {
		limit = len(slice)
	}

	if header {
		fmt.Println("ID\tValue\tCount")
	}

	for _, entry := range slice[:limit] {
		fmt.Printf("%d\t%s\t%d\n", entry.ident, entry.name, entry.count)
	}
}
