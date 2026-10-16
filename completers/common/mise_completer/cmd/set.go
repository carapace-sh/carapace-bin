package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Add/update environment variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().StringP("file", "f", "", "Target config file")
	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
