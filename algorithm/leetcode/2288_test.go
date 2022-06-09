package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	words := strings.Split(sentence, " ")
	for i := range words {
		price, is := isPrice(words[i])
		if is {
			words[i] = fmt.Sprintf("$%.2f", float64(price)*(float64(100-discount)/100.0))
		}
	}
	return strings.Join(words, " ")
}

func isPrice(s string) (int, bool) {
	if len(s) > 1 && s[0] == '$' {
		v, err := strconv.Atoi(s[1:])
		if err == nil {
			return v, true
		}
	}
	return 0, false
}
