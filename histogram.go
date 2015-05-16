package main

import (
	"fmt"
	"sort"
	"github.com/palmergs/tokensearch"
)

type HistoEntry struct {
	name		string
	count		int
}

type HistoEntries []*HistoEntry

func NewHistoEntries(matches []*tokensearch.TokenMatch) HistoEntries {
	histo := make(map[string]*HistoEntry)
	for _, match := range matches {
		key := match.Token.Display
		if histo[key] == nil {
			histo[key] = &HistoEntry{key, 0}
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

func (slice HistoEntries) PrintList(limit int) {
	sort.Sort(slice)
	for _, entry := range slice[:limit] {
		fmt.Printf("%s\t%d\n", entry.name, entry.count)
	}
}