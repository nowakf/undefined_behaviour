package common

import "unicode"

func init() {
	c := cat('█')
	for _, v := range c {
		println(v, "unicode rangetable")

	}
}

func cat(r rune) (names []string) {
	names = make([]string, 0)
	for name, table := range unicode.Categories {
		if unicode.Is(table, r) {
			names = append(names, name)
		}
	}
	return
}
