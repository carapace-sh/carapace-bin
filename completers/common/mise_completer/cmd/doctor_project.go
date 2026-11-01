package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctor_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Run the project's diagnostic checks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctor_projectCmd).Standalone()

	doctor_projectCmd.Flags().BoolP("help", "h", false, "Print help")
	doctor_projectCmd.Flags().BoolP("json", "J", false, "Output the complete report as JSON")
	doctorCmd.AddCommand(doctor_projectCmd)
}
