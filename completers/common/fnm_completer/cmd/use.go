package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Change Node.js version",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(useCmd)

	addCommonFlags(useCmd)

	useCmd.Flags().Bool("install-if-missing", false, "Install the version if it isn't installed yet")
	useCmd.Flags().Bool("silent-if-unchanged", false, "Don't output a message identifying the version being used if it will not change due to execution of this command")

	carapace.Gen(useCmd).Standalone()

	carapace.Gen(useCmd).PositionalAnyCompletion(
		action.ActionInstalledVersions(),
	)
}
