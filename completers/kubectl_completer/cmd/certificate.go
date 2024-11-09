package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certificateCmd = &cobra.Command{
	Use:     "certificate SUBCOMMAND",
	Short:   "Modify certificate resources",
	GroupID: "cluster management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certificateCmd).Standalone()

	rootCmd.AddCommand(certificateCmd)
}
