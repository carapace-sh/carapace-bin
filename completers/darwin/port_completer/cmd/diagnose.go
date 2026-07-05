package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Check for common issues in the user's environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnoseCmd).Standalone()
	diagnoseCmd.Flags().Bool("quiet", false, "Only display warnings and errors")
	rootCmd.AddCommand(diagnoseCmd)
}
