package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var programCmd = &cobra.Command{
	Use:   "program",
	Short: "Program deployment and management commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(programCmd).Standalone()

	programCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(programCmd)
}
