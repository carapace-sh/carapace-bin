package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version [options]",
	Short: "Show the current Terraform version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().BoolS("json", "json", false, "Output the version information as a JSON object.")
	rootCmd.AddCommand(versionCmd)
}
