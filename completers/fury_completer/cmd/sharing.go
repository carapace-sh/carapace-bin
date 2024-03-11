package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharingCmd = &cobra.Command{
	Use:   "sharing",
	Short: "Collaboration commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharingCmd).Standalone()

	rootCmd.AddCommand(sharingCmd)
}
