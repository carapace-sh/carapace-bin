package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getenvCmd = &cobra.Command{
	Use:   "getenv",
	Short: "Get the value of an environment variable from within launchd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getenvCmd).Standalone()
	rootCmd.AddCommand(getenvCmd)
}
