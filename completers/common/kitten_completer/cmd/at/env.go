package at

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Change environment variables seen by future children",
}

func init() {
	envCmd.AddCommand(atCmd)
	carapace.Gen(envCmd).Standalone()

	envCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(envCmd).FlagCompletion(carapace.ActionMap{})
}