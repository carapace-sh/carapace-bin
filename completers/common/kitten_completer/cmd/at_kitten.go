package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kittenCmd = &cobra.Command{
	Use:   "kitten",
	Short: "Run a kitten",
}

func init() {
	kittenCmd.AddCommand(atCmd)
	carapace.Gen(kittenCmd).Standalone()

	kittenCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(kittenCmd).FlagCompletion(carapace.ActionMap{})
}