package mitmproxy

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionAppendableFiles completes files with optional `+` prefix
//
//	path/to/file
//	+path/to/file
func ActionAppendableFiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.CallbackValue, "+") {
			c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "+")
			return carapace.ActionFiles().Invoke(c).Prefix("+").ToA()
		}
		return carapace.ActionFiles()
	})
}
