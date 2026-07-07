package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join strings with a separator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()

	joinCmd.Flags().BoolP("no-empty", "n", false, "exclude empty strings")
	joinCmd.Flags().BoolP("quiet", "q", false, "suppress output")

	rootCmd.AddCommand(joinCmd)
}
