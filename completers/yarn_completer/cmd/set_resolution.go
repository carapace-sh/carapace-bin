package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_resolutionCmd = &cobra.Command{
	Use:   "resolution",
	Short: "Enforce a package resolution",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_resolutionCmd).Standalone()

	set_resolutionCmd.Flags().BoolP("save", "s", false, "Persist the resolution inside the top-level manifest")
	setCmd.AddCommand(set_resolutionCmd)

	// TODO positional completion
}
