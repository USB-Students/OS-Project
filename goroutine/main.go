package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GoID Code copied from https://gist.github.com/metafeather/3615b23097836bc36579100dac376906
// this shouldn't be used in peal projects
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
