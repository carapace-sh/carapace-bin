package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var starsCmd = &cobra.Command{
	Use:   "stars",
	Short: "Lists all packages starred by a specific user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(starsCmd).Standalone()

	starsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(starsCmd)
}
