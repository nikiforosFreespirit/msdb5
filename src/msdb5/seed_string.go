// Code generated by "stringer -type=Seed ./src/msdb5/card.go"; DO NOT EDIT.

package msdb5

import "strconv"

const _Seed_name = "CoinCupSwordCudgel"

var _Seed_index = [...]uint8{0, 4, 7, 12, 18}

func (i Seed) String() string {
	if i >= Seed(len(_Seed_index)-1) {
		return "Seed(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Seed_name[_Seed_index[i]:_Seed_index[i+1]]
}
