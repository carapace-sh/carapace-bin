package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var coCmd = &cobra.Command{
	Use:     "co",
	Short:   "Alias for \"pr checkout\"",
	GroupID: "alias",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coCmd).Standalone()

	rootCmd.AddCommand(coCmd)
}
