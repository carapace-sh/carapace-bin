package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run the doctor to check for common issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_doctorCmd).Standalone()

	helpCmd.AddCommand(help_doctorCmd)
}
