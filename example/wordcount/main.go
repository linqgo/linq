package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/linqgo/linq/v2"
)

func main() {
	text := `To be or not to be that is the question
Whether tis nobler in the mind to suffer
The slings and arrows of outrageous fortune
Or to take arms against a sea of troubles
And by opposing end them to die to sleep
No more and by a sleep to say we end
The heartache and the thousand natural shocks
That flesh is heir to tis a consummation`

	// Split text into lines, then flatten into individual words.
	lines := slices.Values(strings.Split(text, "\n"))
	words := linq.SelectMany(lines, func(line string) iter.Seq[string] {
		return slices.Values(strings.Fields(strings.ToLower(line)))
	})

	// Group by word and count occurrences.
	groups := linq.GroupBySlices(words, func(w string) string { return w })

	type WordCount struct {
		Word  string
		Count int
	}

	// Convert groups to WordCount structs.
	counts := linq.Select(groups, func(g linq.KV[string, []string]) WordCount {
		return WordCount{Word: g.Key, Count: len(g.Value)}
	})

	// Sort by frequency descending, take top 10.
	sorted := linq.OrderByDesc(counts, func(wc WordCount) int { return wc.Count })
	top10 := linq.Take(sorted, 10)

	fmt.Println("Top 10 words by frequency:")
	for wc := range top10 {
		fmt.Printf("  %-12s %d\n", wc.Word, wc.Count)
	}
}
