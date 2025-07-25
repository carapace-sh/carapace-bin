//go:build !windows

package shim

import (
	"fmt"
)

func body(target string) []byte {
	return fmt.Appendf(nil, "#!/bin/sh\nexec carapace --run \"%v\" \"$@\"", target)
}
