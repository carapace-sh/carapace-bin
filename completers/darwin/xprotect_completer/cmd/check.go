package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "print the currently online available update version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().Bool("json", false, "Output in JSON format")
}
