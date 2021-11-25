package rune

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRune(t *testing.T) {
	var a = "龙志亮"
	assert.Equal(t, 3, len([]rune(a)))
	assert.Equal(t, 9, len([]byte(a)))
	assert.Equal(t, "龙", string([]rune(a)[:1]))
	assert.Equal(t, "志亮", string([]rune(a)[1:]))

}
