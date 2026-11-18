package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trickOrTreatCmd = &cobra.Command{
	Use:    "trick-or-treat",
	Short:  "Dig a forgotten command, or a ghost, out of your history",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trickOrTreatCmd).Standalone()

	trickOrTreatCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(trickOrTreatCmd)
}
