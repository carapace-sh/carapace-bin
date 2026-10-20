package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check the health of your Claude Code installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(doctorCmd)
}
