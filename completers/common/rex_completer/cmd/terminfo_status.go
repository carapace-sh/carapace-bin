package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report the TERM a new terminal would be given, and why",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoStatusCmd).Standalone()

	terminfoStatusCmd.Flags().Bool("json", false, "output the report as JSON")

	terminfoCmd.AddCommand(terminfoStatusCmd)
}
