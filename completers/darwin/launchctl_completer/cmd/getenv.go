package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getenvCmd = &cobra.Command{
	Use:   "getenv",
	Short: "Gets the value of an environment variable from within launchd",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(getenvCmd).Standalone()
	rootCmd.AddCommand(getenvCmd)
}
