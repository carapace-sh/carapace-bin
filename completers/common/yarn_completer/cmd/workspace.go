package cmd

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var workspaceCmd = &cobra.Command{
	Use:     "workspace",
	Short:   "Run a command within the specified workspace",
	GroupID: "workspace-related",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspaceCmd).Standalone()

	rootCmd.AddCommand(workspaceCmd)

	carapace.Gen(workspaceCmd).PositionalCompletion(
		yarn.ActionWorkspaces(),
	)

	carapace.Gen(workspaceCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("yarn", "workspaces", "list", "--json")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")

				for _, line := range lines[:len(lines)-1] {
					var workspace struct {
						Name     string
						Location string
					}
					if err := json.Unmarshal([]byte(line), &workspace); err != nil {
						return carapace.ActionMessage(err.Error())
					}

					if workspace.Name == c.Args[0] {
						return bridge.ActionCarapaceBin("yarn").Chdir(workspace.Location).Shift(1)
					}
				}
				return carapace.ActionMessage("unknown workspace: %#v", c.Args[0])
			})
		}),
	)
}
