package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Amend a file change into a specific commit and rebases any dependent commits.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amendCmd).Standalone()

	amendCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(amendCmd)
}
