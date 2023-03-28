package cmd

import (
	"encoding/json"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/yarn"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
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
						c.Dir = workspace.Location
						c.Args = c.Args[1:]
						return bridge.ActionCarapaceBin("yarn").Invoke(c).ToA()
					}
				}
				return carapace.ActionMessage("unknown workspace: %#v", c.Args[0])
			})
		}),
	)
}
