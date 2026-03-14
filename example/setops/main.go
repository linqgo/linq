package main

import (
	"fmt"
	"slices"

	"github.com/linqgo/linq/v2"
)

func main() {
	teamA := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	teamB := []string{"Charlie", "Diana", "Frank", "Grace", "Alice"}

	seqA := slices.Values(teamA)
	seqB := slices.Values(teamB)

	fmt.Println("Team A:", teamA)
	fmt.Println("Team B:", teamB)
	fmt.Println()

	// Union: all unique members across both teams.
	fmt.Println("Union (all unique members):")
	for name := range linq.Union(seqA, seqB) {
		fmt.Printf("  %s\n", name)
	}

	// Intersect: members on both teams.
	fmt.Println("\nIntersect (on both teams):")
	for name := range linq.Intersect(seqA, seqB) {
		fmt.Printf("  %s\n", name)
	}

	// Except: members only on Team A.
	fmt.Println("\nExcept (only on Team A):")
	for name := range linq.Except(seqA, seqB) {
		fmt.Printf("  %s\n", name)
	}

	// Distinct on a list with duplicates.
	withDups := []string{"red", "blue", "red", "green", "blue", "red", "yellow"}
	fmt.Println("\nOriginal colors:", withDups)
	fmt.Print("Distinct colors:")
	for color := range linq.Distinct(slices.Values(withDups)) {
		fmt.Printf(" %s", color)
	}
	fmt.Println()
}
