package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compareVersionsCmd = &cobra.Command{
	Use:   "compare-versions",
	Short: "Compare version numbers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compareVersionsCmd).Standalone()

	carapace.Gen(compareVersionsCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValuesDescribed(
			"lt", "treat empty version as earlier than any version",
			"le", "treat empty version as earlier than any version",
			"eq", "treat empty version as earlier than any version",
			"ne", "treat empty version as earlier than any version",
			"ge", "treat empty version as earlier than any version",
			"gt", "treat empty version as earlier than any version",

			"lt-nl", "treat empty version as later than any version",
			"le-nl", "treat empty version as later than any version",
			"ge-nl", "treat empty version as later than any version",
			"gt-nl", "treat empty version as later than any version",

			"<", "only for compatibility with control file syntax",
			"<<", "only for compatibility with control file syntax",
			"<=", "only for compatibility with control file syntax",
			"=", "only for compatibility with control file syntax",
			">=", "only for compatibility with control file syntax",
			">>", "only for compatibility with control file syntax",
			">", "only for compatibility with control file syntax",
		),
		carapace.ActionValues(),
	)
}
