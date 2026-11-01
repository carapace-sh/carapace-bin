package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Show ASN of applications matching key=value pairs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findCmd).Standalone()
	findCmd.Flags().Bool("includeExitedApplications", false, "Include exited applications")
	rootCmd.AddCommand(findCmd)
}
