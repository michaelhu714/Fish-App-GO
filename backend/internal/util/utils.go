package util

import (
	"github.com/michaelhu714/Fish-App-GO/types"
)

func intersection (map[types.Card]types.Card s1, map[types.Card]types.Card s2) map[types.Card]types.Card {
	var inter map[types.Card]types.Card
	for k, v := range s1 {
		if val, ok := s2[key]; ok {
			iner[key] = val
		}
	}
	return inter
}
