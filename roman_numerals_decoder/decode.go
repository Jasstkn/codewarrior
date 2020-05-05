package roman_numerals_decoder

type Decoded interface {
	Decode() int
}

type Roman struct {
	Phrase string
}

func (r Roman) Decode() int {
	var result int

	alphabet := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for pos, char := range r.Phrase {
		if pos < len(r.Phrase)-1 && alphabet[string(r.Phrase[pos])] < alphabet[string(r.Phrase[pos+1])] {
			result -= alphabet[string(char)]
		} else {
			result += alphabet[string(char)]
		}
	}
	return result
}
