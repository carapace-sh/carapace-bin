package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:     "doctor",
	Short:   "Check your mise setup for problems",
	Aliases: []string{"dr"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctorCmd).Standalone()

	doctorCmd.Flags().BoolP("help", "h", false, "Print help")
	doctorCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	rootCmd.AddCommand(doctorCmd)
}
