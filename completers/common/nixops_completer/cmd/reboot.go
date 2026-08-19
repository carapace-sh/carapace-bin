package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RebootCmd).Standalone()
	rootCmd.AddCommand(RebootCmd)
}
