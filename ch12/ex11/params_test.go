package params

import "testing"

func TestPack(t *testing.T) {
	type Params struct {
		Language string `http:"l"`
		Max      int    `http:"max"`
	}
	p := Params{
		"golang",
		100,
	}
	expect := "l=golang&max=100"
	u, err := Pack(&p)
	if err != nil {
		t.Errorf("Pack(%#v): %s", p, err)
	}
	actual := u.RawQuery
	if actual != expect {
		t.Errorf("Pack(%#v): actual=%q, expect=%q", p, actual, expect)
	}
}
