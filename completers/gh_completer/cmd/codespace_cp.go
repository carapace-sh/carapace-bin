package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var codespace_cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy files between local and remote file systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	codespace_cpCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_cpCmd.Flags().BoolP("expand", "e", false, "Expand remote file names on remote shell")
	codespace_cpCmd.Flags().BoolP("recursive", "r", false, "Recursively copy directories")
	codespaceCmd.AddCommand(codespace_cpCmd)

	carapace.Gen(codespace_cpCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
	})

	carapace.Gen(codespace_cpCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "remote:")
			return action.ActionCodespacePath(
				codespace_cpCmd.Flag("codespace").Value.String(),
				codespace_cpCmd.Flag("expand").Changed,
			).Invoke(c).Prefix("remote:").ToA()
		}),
	)
}
