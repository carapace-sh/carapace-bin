package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Shows a short information about Composer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aboutCmd).Standalone()

	rootCmd.AddCommand(aboutCmd)
}
