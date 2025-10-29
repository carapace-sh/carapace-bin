package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "The default components installed with a toolchain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_profileCmd).Standalone()

	set_profileCmd.Flags().BoolP("help", "h", false, "Print help")
	setCmd.AddCommand(set_profileCmd)

	carapace.Gen(set_profileCmd).PositionalCompletion(
		carapace.ActionValues("minimal", "default", "complete"),
	)
}
