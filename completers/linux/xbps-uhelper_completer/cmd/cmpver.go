package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cmpverCommand = &cobra.Command{
	Use:   "cmpver <version> <version>",
	Short: "Compare two version strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cmpverCommand).Standalone()

	rootCmd.AddCommand(cmpverCommand)
}
