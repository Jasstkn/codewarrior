package roman_numerals_decoder

import (
	"testing"
)

func TestDecode(t *testing.T) {
	romanTests := []struct {
		decoded Decoded
		result  int
	}{
		{decoded: Roman{"XXI"}, result: 21},
		{decoded: Roman{"I"}, result: 1},
		{decoded: Roman{"IV"}, result: 4},
		{decoded: Roman{"MMVIII"}, result: 2008},
		{decoded: Roman{"MDCLXVI"}, result: 1666},

	}

	for _, tt := range romanTests {
		got := tt.decoded.Decode()
		if got != tt.result {
			t.Errorf("%#v got %d want %d", tt.decoded, got, tt.result)
		}
	}
}
