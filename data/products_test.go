package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Lol",
		Price: 1,
		SKU:   "wed-awer-asfe",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
