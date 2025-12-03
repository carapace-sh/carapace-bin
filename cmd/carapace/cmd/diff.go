package cmd

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/action"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	shlex "github.com/carapace-sh/carapace-shlex"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "--diff COMPLETER COMPLETER ...ARGS",
	Short: "diff completion",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		lines := carapace.DiffPatch(
			bridge.ActionCarapaceBin(args[0]),
			bridge.ActionCarapaceBin(args[1]),
			carapace.NewContext(args[2:]...),
		)

		fmt.Printf("--- a/carapace %v\n", shlex.Join(append([]string{args[0], "export"}, args[2:]...)))
		fmt.Printf("+++ b/carapace %v\n", shlex.Join(append([]string{args[1], "export"}, args[2:]...)))
		fmt.Println(strings.Join(lines, "\n"))
	},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
	diffCmd.Flags().SetInterspersed(false)

	carapace.Gen(diffCmd).PositionalCompletion(
		action.ActionCompleters(),
		action.ActionCompleters(),
	)

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Diff(
				bridge.ActionCarapaceBin(c.Args[0]).Shift(2),
				bridge.ActionCarapaceBin(c.Args[1]).Shift(2),
			)
		}),
	)
}
