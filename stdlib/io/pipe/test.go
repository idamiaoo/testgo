package pipe

import (
	"io"
	"testing"
	"time"
)

func TestPipe(t *testing.T) {
	io.Pipe()
	time.Now()
}
