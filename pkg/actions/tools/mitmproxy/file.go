package mitmproxy

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionAppendableFiles completes files with optional `+` prefix
//
//	path/to/file
//	+path/to/file
func ActionAppendableFiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if after, ok := strings.CutPrefix(c.Value, "+"); ok {
			c.Value = after
			return carapace.ActionFiles().Invoke(c).Prefix("+").ToA()
		}
		return carapace.ActionFiles()
	})
}
