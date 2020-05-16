package cmd

import (
	"github.com/rsteube/carapace"
)

func init() {
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":      carapace.ActionValues("always", "auto", "never"),
		"colour":     carapace.ActionValues("always", "auto", "never"),
		"sort":       carapace.ActionValues("name", "Name", "size", "extension", "Extension", "modified", "changed", "accessed", "created", "inode", "type", "none"),
		"time":       carapace.ActionValues("modified", "accessed", "created"),
		"time-style": carapace.ActionValues("default", "iso", "long-iso", "full-iso"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(""),
	)
}
