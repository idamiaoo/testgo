package pipe

import (
	"io"
	"testing"
	"time"
)

func TestPipe(t *testing.T) {
	io.Pipe()
	io.TeeReader()
	time.Now()
}
