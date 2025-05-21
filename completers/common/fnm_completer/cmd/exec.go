package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/fnm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command within fnm context",
	Run:   func(*cobra.Command, []string) {},
}

func init() {
	rootCmd.AddCommand(execCmd)

	addCommonFlags(execCmd)

	execCmd.Flags().String("using", "", "Either an explicit version, or a filename with the version written in it")

	carapace.Gen(execCmd).Standalone()

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"using": action.ActionLocalVersions(),
	})
}
