package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bumpCmd = &cobra.Command{
	Use:   "bump",
	Short: "Update checksums and revision for a new version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpCmd).Standalone()
	rootCmd.AddCommand(bumpCmd)
}
