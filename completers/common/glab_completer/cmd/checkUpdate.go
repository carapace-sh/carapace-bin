package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkUpdateCmd = &cobra.Command{
	Use:     "check-update",
	Short:   "Check for the latest glab version.",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkUpdateCmd).Standalone()

	rootCmd.AddCommand(checkUpdateCmd)
}
