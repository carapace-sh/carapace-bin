//go:build !windows

package shim

import (
	"fmt"

	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func body(target string) []byte {
	return fmt.Appendf(nil, "#!/bin/sh\nexec carapace --run \"%v\" \"$@\"", shlex.Escape(target, shlex.Bash))
}
