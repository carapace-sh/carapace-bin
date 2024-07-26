package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eventGenerateCmd = &cobra.Command{
	Use:   "event:generate",
	Short: "Generate the missing events and listeners based on registration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventGenerateCmd).Standalone()

	rootCmd.AddCommand(eventGenerateCmd)
}
