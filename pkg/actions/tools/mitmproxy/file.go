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
		if strings.HasPrefix(c.Value, "+") {
			c.Value = strings.TrimPrefix(c.Value, "+")
			return carapace.ActionFiles().Invoke(c).Prefix("+").ToA()
		}
		return carapace.ActionFiles()
	})
}
