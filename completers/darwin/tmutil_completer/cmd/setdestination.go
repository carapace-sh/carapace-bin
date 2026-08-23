package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setdestinationCmd = &cobra.Command{
	Use:   "setdestination",
	Short: "configure a backup destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setdestinationCmd).Standalone()
	rootCmd.AddCommand(setdestinationCmd)

	setdestinationCmd.Flags().BoolP("add", "a", false, "Add to the list of destinations")
	setdestinationCmd.Flags().BoolP("password", "p", false, "Enter password at interactive prompt")

	carapace.Gen(setdestinationCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
