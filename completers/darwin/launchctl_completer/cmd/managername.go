package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managernameCmd = &cobra.Command{
	Use:   "managername",
	Short: "Print the name of the current launchd session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managernameCmd).Standalone()
	rootCmd.AddCommand(managernameCmd)
}
