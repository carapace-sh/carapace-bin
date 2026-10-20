// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var doctor_pathCmd = &cobra.Command{
	Use:     "path",
	Short:   "Print the directories mise adds to PATH",
	Aliases: []string{"paths"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(doctor_pathCmd).Standalone()

	doctor_pathCmd.Flags().BoolP("full", "f", false, "Print every PATH entry, including those mise does not provide")
	doctor_pathCmd.Flags().BoolP("help", "h", false, "Print help")
	doctorCmd.AddCommand(doctor_pathCmd)
}
