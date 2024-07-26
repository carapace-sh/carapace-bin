package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeEventCmd = &cobra.Command{
	Use:   "make:event [-f|--force] [--] <name>",
	Short: "Create a new event class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeEventCmd).Standalone()

	makeEventCmd.Flags().Bool("force", false, "Create the class even if the event already exists")
	rootCmd.AddCommand(makeEventCmd)
}
