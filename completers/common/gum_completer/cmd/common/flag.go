package common

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func AddFlagCompletion(cmd *cobra.Command) {
	m := make(carapace.ActionMap)
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		switch strings.TrimPrefix(filepath.Ext(f.Name), ".") {
		case "background", "foreground", "border-background", "border-foreground":
			m[f.Name] = gum.ActionColors()
		case "align":
			m[f.Name] = gum.ActionAlignments()
		case "border":
			m[f.Name] = gum.ActionBorders()
		}
	})

	carapace.Gen(cmd).FlagCompletion(m)
}
