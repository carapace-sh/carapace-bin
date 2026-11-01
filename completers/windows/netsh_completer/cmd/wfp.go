package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wfpCmd = &cobra.Command{
	Use:   "wfp",
	Short: "Windows Filtering Platform configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wfpCmd).Standalone()
	rootCmd.AddCommand(wfpCmd)
}
