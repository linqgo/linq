package main

import (
	"fmt"
	"slices"

	"github.com/linqgo/linq/v2"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	people := []Person{
		{"Alice", 32, "London"},
		{"Bob", 17, "Paris"},
		{"Charlie", 45, "London"},
		{"Diana", 28, "Berlin"},
		{"Eve", 15, "Paris"},
		{"Frank", 52, "Berlin"},
		{"Grace", 22, "London"},
		{"Hank", 19, "Paris"},
	}

	// Filter to adults only, then group by city.
	adults := linq.Where(slices.Values(people), func(p Person) bool {
		return p.Age > 18
	})

	groups := linq.GroupBySlices(adults, func(p Person) string {
		return p.City
	})

	// For each city, count people and find the oldest.
	for group := range groups {
		city := group.Key
		members := group.Value

		count := len(members)

		oldest := members[0]
		for _, p := range members[1:] {
			if p.Age > oldest.Age {
				oldest = p
			}
		}

		fmt.Printf("%s: %d adults, oldest is %s (age %d)\n",
			city, count, oldest.Name, oldest.Age)
	}
}
