package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revisionCmd = &cobra.Command{
	Use:   "revision",
	Short: "Revision commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revisionCmd).Standalone()

	revisionCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(revisionCmd)
}
