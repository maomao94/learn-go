package container

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"pwwwwkew", 4},
		{"", 0},
		{"a", 1},
		{"b", 1},
		{"种种阿", 1},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s;"+
				"expexted %d", actual, tt.s, tt.ans)
		}
	}

}

func BenchmarkSubstr(b *testing.B) {
	s := "jkdajkdkadkdjkdjkadhjkjak"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s;"+
				"expexted %d", actual, s, ans)
		}
	}
}
