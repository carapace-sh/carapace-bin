package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_installmanpagesCmd = &cobra.Command{
	Use:   "install-man-pages",
	Short: "Install Jujutsu's manpages to the provided path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_installmanpagesCmd).Standalone()

	util_installmanpagesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")

	carapace.Gen(util_installmanpagesCmd).PositionalCompletion(carapace.ActionFiles())

	utilCmd.AddCommand(util_installmanpagesCmd)
}
