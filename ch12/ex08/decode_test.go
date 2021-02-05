package decode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	s := "((Title \"Dr. Strangelove\") (Subtitle \"How I Learned to Stop Worrying and Love the Bomb\") (Year 1964) (Color nil) (Actor ((\"Dr. Strangelove\" \"Peter Sellers\") (\"Grp. Capt. Lionel Mandrake\" \"Peter Sellers\") (\"Pres. Merkin Muffley\" \"Peter Sellers\") (\"Gen. Buck Turgidson\" \"George C. Scott\") (\"Brig. Gen. Jack D. Ripper\" \"Sterling Hayden\"))) (Oscars (\"Best Actor (Nomin.)\" \"Best Adapted Screenplay (Nomin.)\" \"Best Director (Nomin.)\" \"Best Picture (Nomin.)\")) (Sequel nil))"
	b := []byte(s)
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	var data Movie
	if err := Unmarshal(b, &data); err != nil {
		fmt.Errorf("Marshal(strangelove): %s", err)
	}
	if !reflect.DeepEqual(data, strangelove) {
		t.Fatal("not equal")
	}
}
