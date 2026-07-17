package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var manageruidCmd = &cobra.Command{
	Use:   "manageruid",
	Short: "Print the UID of the current launchd session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(manageruidCmd).Standalone()
	rootCmd.AddCommand(manageruidCmd)
}
