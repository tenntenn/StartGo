package pairing

import (
	"math/rand"
)

func randomSelect(rnd *rand.Rand, strs []string) (string, []string) {
	i := rnd.Intn(len(strs))
	return strs[i], append(strs[:i], strs[i+1:]...)
}

func MakePair(rnd *rand.Rand, strs []string) [][]string {
	pairs := make([][]string, 0)
	cpy := make([]string, len(strs))
	copy(cpy, strs)
	for len(cpy) > 1 {
		var str1, str2 string
		str1, cpy = randomSelect(rnd, cpy)
		str2, cpy = randomSelect(rnd, cpy)
		pair := []string{str1, str2}
		pairs = append(pairs, pair)
	}

	return pairs
}
