package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set SUBCOMMAND",
	Short:   "Set specific features on objects",
	GroupID: "basic beginner",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	rootCmd.AddCommand(setCmd)
}
