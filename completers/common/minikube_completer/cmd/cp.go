package cmd

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:     "cp",
	Short:   "Copy the specified file into minikube",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()
	rootCmd.AddCommand(cpCmd)

	carapace.Gen(cpCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action { // TODO this won't support paths containing `:` in minikube
			switch len(c.Parts) {
			case 0:
				return action.ActionNodes().Invoke(c).Suffix(":/").ToA()
			case 1:
				node := c.Parts[0]
				path := filepath.Dir(c.Value)
				if path == "" {
					path = "/"
				}
				return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
					return carapace.ActionExecCommand("minikube", "ssh", "--node", node, "ls", `\-1`, `\-p`, path)(func(output []byte) carapace.Action {
						lines := strings.Split(string(output), "\n")
						return carapace.ActionValues(lines[:len(lines)-1]...)
					})
				})
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
