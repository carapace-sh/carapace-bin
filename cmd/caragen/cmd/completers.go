package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completersCmd = &cobra.Command{
	Use:   "completers",
	Short: "list completers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completersCmd).Standalone()

	rootCmd.AddCommand(completersCmd)

	carapace.Gen(completersCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"common", "all ${GOOS}",
			"darwin", "common and macOS",
			"linux", "common and linux",
			"unix", "common, linux and macOS",
			"windows", "common and windows",
		),
	)
}
