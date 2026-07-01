package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managerpidCmd = &cobra.Command{
	Use:   "managerpid",
	Short: "Prints the PID of the launchd controlling the session",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(managerpidCmd).Standalone()
	rootCmd.AddCommand(managerpidCmd)
}
