package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Verify package installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().StringP("component", "c", "", "Check installed packages under given component")
	checkCmd.Flags().Bool("config", false, "Check only changed config files of the packages")

	rootCmd.AddCommand(checkCmd)
}
