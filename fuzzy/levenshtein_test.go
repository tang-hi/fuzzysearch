package fuzzy

import (
	"crypto/rand"
	random "math/rand"
	"testing"
)

var levenshteinDistanceTests = []struct {
	s, t   string
	wanted int
}{
	{"zazz", deBelloGallico + " zazz", 1544},
	{"zazz", "zazz " + deBelloGallico, 1544},
	{"a", "a", 0},
	{"ab", "ab", 0},
	{"ab", "aa", 1},
	{"ab", "aaa", 2},
	{"bbb", "a", 3},
	{"kitten", "sitting", 3},
	{"ёлка", "ёлочка", 2},
	{"ветер", "ёлочка", 6},
	{"中国", "中华人民共和国", 5},
	{"日本", "中华人民共和国", 7},
}

func TestLevenshtein(t *testing.T) {
	for _, test := range levenshteinDistanceTests {
		distance := OSADamerauLevenshteinDistance(test.s, test.t)
		if distance != test.wanted {
			t.Errorf("got distance %d, expected %d for %s in %s",
				distance, test.wanted, test.s, test.t)
		}
	}
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LevenshteinDistance(randString(random.Intn(500)), randString(random.Intn(500)))
	}
}

// func BenchmarkLevenshteinDistanceBigLate(b *testing.B) {
// 	ldt := levenshteinDistanceTests[0]
// 	for i := 0; i < b.N; i++ {
// 		LevenshteinDistance(ldt.s, ldt.t)
// 	}
// }

// func BenchmarkLevenshteinDistanceBigEarly(b *testing.B) {
// 	ldt := levenshteinDistanceTests[1]
// 	for i := 0; i < b.N; i++ {
// 		LevenshteinDistance(ldt.s, ldt.t)
// 	}
// }

func BenchmarkOSDLevenshteinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OSADamerauLevenshteinDistance(randString(random.Intn(500)), randString(random.Intn(500)))
	}
}

// func BenchmarkOSDLevenshteinDistanceBigLate(b *testing.B) {
// 	ldt := levenshteinDistanceTests[0]
// 	for i := 0; i < b.N; i++ {
// 		OSADamerauLevenshteinDistance(ldt.s, ldt.t)
// 	}
// }

// func BenchmarkOSDLevenshteinDistanceBigEarly(b *testing.B) {
// 	ldt := levenshteinDistanceTests[1]
// 	for i := 0; i < b.N; i++ {
// 		OSADamerauLevenshteinDistance(ldt.s, ldt.t)
// 	}
// }

func randString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}
