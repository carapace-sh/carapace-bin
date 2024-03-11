package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whoamiCmd = &cobra.Command{
	Use:     "whoami",
	Short:   "Show current logged in user",
	GroupID: "MISCELLANEOUS",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whoamiCmd).Standalone()

	rootCmd.AddCommand(whoamiCmd)
}
