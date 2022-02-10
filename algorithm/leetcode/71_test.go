package leetcode

import (
	"strings"
	"testing"
)

func simplifyPath(path string) string {
	var index = -1
	var paths []string
	for _, v := range strings.Split(path, "/") {
		switch v {
		case "..":
			if index >= 0 {
				index--
			}
		case "", ".":
		default:
			index++
			if index >= len(paths) {
				paths = append(paths, v)
			} else {
				paths[index] = v
			}
		}
	}
	return "/" + strings.Join(paths[:index+1], "/")
}

/*
"/a/./b/../../c/"
*/

func Test71(t *testing.T) {
	v := strings.Split("/home/", "/")
	t.Log(v, len(v))
	a := strings.Join([]string{}, "/")
	t.Log(a, len(a))
	t.Log(simplifyPath("/a/./b/../../c/"))
}
