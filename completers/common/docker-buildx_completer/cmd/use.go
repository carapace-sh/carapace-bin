package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/docker-buildx_completer/cmd/action"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [OPTIONS] NAME",
	Short: "Set the current builder instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(useCmd).Standalone()
	useCmd.Flags().Bool("default", false, "Set builder as default for current context")
	useCmd.Flags().Bool("global", false, "Builder persists context changes")
	rootCmd.AddCommand(useCmd)

	carapace.Gen(useCmd).PositionalCompletion(
		action.ActionBuilders(),
	)
}
