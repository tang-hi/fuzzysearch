package fuzzy

// LevenshteinDistance measures the difference between two strings.
// The Levenshtein distance between two words is the minimum number of
// single-character edits (i.e. insertions, deletions or substitutions)
// required to change one word into the other.
//
// This implemention is optimized to use O(min(m,n)) space and is based on the
// optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
func LevenshteinDistance(s, t string) int {
	r1, r2 := []rune(s), []rune(t)
	column := make([]int, 1, 64)

	for y := 1; y <= len(r1); y++ {
		column = append(column, y)
	}

	for x := 1; x <= len(r2); x++ {
		column[0] = x

		for y, lastDiag := 1, x-1; y <= len(r1); y++ {
			oldDiag := column[y]
			cost := 0
			if r1[y-1] != r2[x-1] {
				cost = 1
			}
			column[y] = min(column[y]+1, min(column[y-1]+1, lastDiag+cost))
			lastDiag = oldDiag
		}
	}

	return column[len(r1)]
}

func OSADamerauLevenshteinDistance(s, t string) int {
	r1, r2 := []rune(s), []rune(t)
	if len(r1) < len(r2) {
		return OSADamerauLevenshteinDistance(t, s)
	}
	row := min(len(r1)+1, 3)
	matrix := make([][]int, 0, 3)

	for i := 0; i < row; i++ {
		matrix = append(matrix, make([]int, len(r2)+1))
		for j := 0; j <= len(r2); j++ {
			matrix[i][j] = 0
		}
		matrix[i][0] = i
	}

	for j := 0; j <= len(r2); j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= len(r1); i++ {
		matrix[convert(i)][0] = i
		for j := 1; j <= len(r2); j++ {
			cost := 0
			if r1[i-1] != r2[j-1] {
				cost = 1
			}

			matrix[convert(i)][j] = min(matrix[convert(i-1)][j]+1,
				min(matrix[convert(i)][j-1]+1, matrix[convert(i-1)][j-1]+cost))

			if i > 1 && j > 1 && r1[i-1] == r2[j-2] && r1[i-2] == r2[j-1] {
				matrix[convert(i)][j] = min(matrix[convert(i)][j], matrix[convert(i-2)][j-2]+1)
			}
		}
	}
	return matrix[convert(len(r1))][len(r2)]
}

func convert(a int) int {
	return a % 3
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
