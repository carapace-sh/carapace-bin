package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Clear an environment variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unsetCmd).Standalone()

	unsetCmd.Flags().StringP("file", "f", "", "Target config file")
	rootCmd.AddCommand(unsetCmd)

	carapace.Gen(unsetCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
