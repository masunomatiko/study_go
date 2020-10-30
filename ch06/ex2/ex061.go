package intset

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Len() int {
	cnt := 0
	for _, word := range s.words {
		cnt += popcount(word)
	}
	return cnt
}

func popcount(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

func (s *IntSet) Remove(x int) {
	copy(s.words[x:], s.words[x+1:])
	s.words = s.words[:len(s.words)-1]
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = s.words
	return &t
}
