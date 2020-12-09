package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copy files and directories to and from containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()

	cpCmd.Flags().StringP("container", "c", "", "Container name. If omitted, the first container in the pod will be chosen")
	cpCmd.Flags().Bool("no-preserve", false, "The copied file/directory's ownership and permissions will not be preserved in the container")
	rootCmd.AddCommand(cpCmd)

	carapace.Gen(cpCmd).FlagCompletion(carapace.ActionMap{
		"container": carapace.ActionCallback(func(args []string) carapace.Action {
			resource := ""
			if len(args) > 0 && !isFileCompletion(args[0]) {
				resource = args[0]
			} else if len(args) > 1 && !isFileCompletion(args[1]) {
				resource = args[1]
			}
			splitted := strings.SplitN(strings.SplitN(resource, ":", 2)[0], "/", 2)

			if len(splitted) == 2 {
				return action.ActionContainers(splitted[0], "pods/"+splitted[1])
			} else {
				return carapace.ActionMessage("missing resource argument")
			}
		}),
	})

	carapace.Gen(cpCmd).PositionalCompletion(
		ActionPathOrContainer(),
		ActionPathOrContainer(),
	)
}

func ActionPathOrContainer() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if isFileCompletion(carapace.CallbackValue) {
			return carapace.ActionFiles("")
		} else {
			return carapace.ActionMultiParts("/", func(args, parts []string) carapace.Action {
				switch len(parts) {
				case 0:
					return action.ActionResources("", "namespaces").Invoke(args).Suffix("/").ToA()
				case 1:
					return action.ActionResources(parts[0], "pods").Invoke(args).Suffix(":").ToA()
				default:
					return carapace.ActionValues()
				}
			})
		}
	})
}

func isFileCompletion(val string) bool {
	return strings.HasPrefix(val, "/") || strings.HasPrefix(val, ".")
}
