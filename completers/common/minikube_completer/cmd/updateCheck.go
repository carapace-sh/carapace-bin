package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCheckCmd = &cobra.Command{
	Use:     "update-check",
	Short:   "Print current and latest version number",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCheckCmd).Standalone()
	rootCmd.AddCommand(updateCheckCmd)
}
