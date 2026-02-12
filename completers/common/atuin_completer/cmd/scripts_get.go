package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/atuin"
	"github.com/spf13/cobra"
)

var scripts_getCmd = &cobra.Command{
	Use:   "get",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scripts_getCmd).Standalone()

	scripts_getCmd.Flags().BoolP("help", "h", false, "Print help")
	scripts_getCmd.Flags().BoolP("script", "s", false, "Display only the executable script with shebang")
	scriptsCmd.AddCommand(scripts_getCmd)

	carapace.Gen(scripts_getCmd).PositionalCompletion(
		atuin.ActionScripts(),
	)
}
