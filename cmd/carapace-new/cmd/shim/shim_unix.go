//go:build !windows

package shim

import (
	"fmt"
)

func body(target string) []byte {
	return []byte(fmt.Sprintf("#!/bin/sh\ncarapace --run \"%v\" \"$@\"", target))
}
