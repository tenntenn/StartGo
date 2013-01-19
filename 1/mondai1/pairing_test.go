package pairing

import (
	"time"
	"math/rand"
	"fmt"
	"testing"
)

func indexOf(str string, strs []string) int {
	for i, s := range strs {
		if str == s {
			return i
		}
	}

	return -1
}

func TestMakePair(t *testing.T) {
	strs := []string{"A", "B", "C", "D", "E", "F"}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 5回分のサンプルを作成
	pairsSet := make([][][]string, 0)
	duplication := make(map[string]int) 
	for i := 0; i < 5; i++ {
		pairs := MakePair(rnd, strs)
		key := fmt.Sprintf("%v", pairs)
		if count, ok := duplication[key]; ok {
			duplication[key] = count + 1
		} else {
			duplication[key] = 0
		}
		pairsSet = append(pairsSet, pairs)
	}

	// 重複の数
	duplicationCount := len(strs) - len(duplication)
	if duplicationCount >= 2 {
		t.Errorf("ランダムの割には重複が多すぎます")
	}

	// 確かめる
	for _, pairs := range pairsSet {
		// ペアは3つできる
		if len(pairs) != len(strs) / 2 {
			t.Errorf("ペアの数は3です。")
		}

		// 同じ文字が2度出てきてないか？
		// ペアかどうか？
		cpy := make([]string, len(strs))
		copy(cpy, strs)
		for _, pair := range pairs {
			// ペアかどうか？
			if len(pair) != 2 {
				t.Errorf("ペアになっていません。")
			}

			for _, str := range pair {
				i := indexOf(str, cpy)
				if i < 0 {
					t.Errorf("元の文字列スライスに含まれていない文字が含まれているか、同じ文字が使われています。")
				}
				cpy = append(cpy[:i], cpy[i+1:]...)
			}
		}

		// 奇数個の場合は1つ余る
		if (len(strs) % 2 == 0 && len(cpy) != 0) || (len(strs) % 2 != 0 && len(cpy) != 1) {
			t.Errorf("すべての文字が使われていません。")
		}
	}
}
