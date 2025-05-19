package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getDefaultCmd = &cobra.Command{
	Use:     "get-default",
	Short:   "Get the name of the default target",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getDefaultCmd).Standalone()

	rootCmd.AddCommand(getDefaultCmd)
}
