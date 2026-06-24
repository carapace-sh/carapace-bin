package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managernameCmd = &cobra.Command{
	Use:   "managername",
	Short: "Prints the name of the current launchd session",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(managernameCmd).Standalone()
	rootCmd.AddCommand(managernameCmd)
}
