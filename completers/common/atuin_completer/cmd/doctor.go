package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run the doctor to check for common issues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(doctorCmd)
}
