package linq_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcelocantos/linq"
)

func TestGroupJoin(t *testing.T) {
	t.Parallel()

	type Person struct {
		Name string
	}

	type Pet struct {
		Name  string
		Owner Person
	}

	type Ownership struct {
		Owner string
		Pets  []string
	}

	magnus := Person{Name: "Hedlund, Magnus"}
	terry := Person{Name: "Adams, Terry"}
	charlotte := Person{Name: "Weiss, Charlotte"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	people := linq.From(magnus, terry, charlotte)
	pets := linq.From(barley, boots, whiskers, daisy)

	// Create a list where each element is an anonymous
	// type that contains a person's name and
	// a collection of names of the pets they own.
	query := linq.GroupJoin(people, pets,
		linq.Identity[Person],
		func(pet Pet) Person { return pet.Owner },
		func(person Person, pets linq.Query[Pet]) Ownership {
			return Ownership{
				Owner: person.Name,
				Pets:  linq.Select(pets, func(pet Pet) string { return pet.Name }).ToSlice(),
			}
		},
	)

	assert.Equal(t,
		[]Ownership{
			{"Hedlund, Magnus", []string{"Daisy"}},
			{"Adams, Terry", []string{"Barley", "Boots"}},
			{"Weiss, Charlotte", []string{"Whiskers"}},
		},
		query.ToSlice(),
	)
}
