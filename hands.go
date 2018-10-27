package main

func onePair(cards []Card) bool {
	comb := combinations(cards, 2, 10)
	count := 0
	for com := range comb {
		if com[0].Number == com[1].Number {
			count++
		}
	}
	return count == 1
}

func twoPair(cards []Card) bool {
	comb := combinations(cards, 2, 10)
	set := [][]Card{}
	for com := range comb {
		if com[0].Number == com[1].Number {
			set = append(set, []Card{com[0], com[1]})
		}
	}
	if len(set) != 2 {
		return false
	}
	return duplicatePair(set[0], set[1])
}

func combinations(list []Card, select_num, buf int) (c chan []Card) {
	c = make(chan []Card, buf)
	go func() {
		defer close(c)
		switch {
		case select_num == 0:
			c <- []Card{}
		case select_num == len(list):
			c <- list
		case len(list) < select_num:
			return
		default:
			for i := 0; i < len(list); i++ {
				for sub_comb := range combinations(list[i+1:], select_num-1, buf) {
					c <- append([]Card{list[i]}, sub_comb...)
				}
			}
		}
	}()
	return
}

func duplicatePair(a []Card, b []Card) bool {
	if a[0].Same(b[0]) {
		return false
	}
	if a[0].Same(b[1]) {
		return false
	}
	if a[1].Same(b[0]) {
		return false
	}
	if a[1].Same(b[1]) {
		return false
	}
	return true
}
