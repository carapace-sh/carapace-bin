package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_backend_nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Print the name of the backend used in the current repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_backend_nameCmd).Standalone()

	util_backend_nameCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	util_backendCmd.AddCommand(util_backend_nameCmd)
}
