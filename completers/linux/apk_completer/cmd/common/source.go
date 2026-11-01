package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddSourceFlags(cmd *cobra.Command) {
	cmd.Flags().String("from", "", "search packages from given source")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"from": carapace.ActionValuesDescribed(
			"installed", "exclude normal repositories",
			"none", "commandline repositories only",
			"repositories", "exclude installed database",
			"system", "all system sources",
		),
	})

}
