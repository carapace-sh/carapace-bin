package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var specsCmd = &cobra.Command{
	Use:   "specs",
	Short: "manage specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(specsCmd).Standalone()

	specsCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.AddCommand(specsCmd)
}
