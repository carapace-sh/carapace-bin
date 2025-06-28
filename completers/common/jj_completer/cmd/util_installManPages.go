package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_installManPagesCmd = &cobra.Command{
	Use:   "install-man-pages",
	Short: "Install Jujutsu's manpages to the provided path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_installManPagesCmd).Standalone()

	util_installManPagesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_installManPagesCmd)

	carapace.Gen(util_installManPagesCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
