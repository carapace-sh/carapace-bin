package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var examineCmd = &cobra.Command{
	Use:   "examine",
	Short: "Run the specified analysis tool against launchd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(examineCmd).Standalone()
	rootCmd.AddCommand(examineCmd)
	carapace.Gen(examineCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
