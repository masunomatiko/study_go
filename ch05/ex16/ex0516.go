package main

func join(elems ...string) string {
	out := ""
	for _, e := range elems {
		out += e
	}
	return out
}
