package lib

import (
	"runtime"
	"testing"
)

func TestDrawJpg(t *testing.T) {
	path := `/Users/v_yangsen08/Downloads/01.webp`
	color := `red`
	DrawWebp(path, color)
	t.Log(runtime.GOOS)
}
