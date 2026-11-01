package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var formatterCmd = &cobra.Command{
	Use:     "formatter",
	Short:   "build or run the formatter",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatterCmd).Standalone()

	rootCmd.AddCommand(formatterCmd)
}
