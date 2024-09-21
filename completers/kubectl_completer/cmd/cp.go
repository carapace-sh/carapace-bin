package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:     "cp <file-spec-src> <file-spec-dest>",
	Short:   "Copy files and directories to and from containers",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()

	cpCmd.Flags().StringP("container", "c", "", "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen")
	cpCmd.Flags().Bool("no-preserve", false, "The copied file/directory's ownership and permissions will not be preserved in the container")
	cpCmd.Flags().String("retries", "", "Set number of retries to complete a copy operation from a container. Specify 0 to disable or any negative value for infinite retrying. The default is 0 (no retry).")
	rootCmd.AddCommand(cpCmd)

	carapace.Gen(cpCmd).FlagCompletion(carapace.ActionMap{
		"container": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			resource := ""
			if len(c.Args) > 0 && !condition.File(c.Args[0])(c) {
				resource = c.Args[0]
			} else if len(c.Args) > 1 && !condition.File(c.Args[1])(c) {
				resource = c.Args[1]
			}
			splitted := strings.SplitN(strings.SplitN(resource, ":", 2)[0], "/", 2)

			if len(splitted) == 2 {
				return kubectl.ActionContainers(kubectl.ContainerOpts{
					Namespace: rootCmd.Flag("namespace").Value.String(),
					Resource:  "pods/" + splitted[1],
				})
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
	return carapace.Batch(
		carapace.ActionFiles(),
		carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return kubectl.ActionResources(kubectl.ResourceOpts{
					Context:   rootCmd.Flag("context").Value.String(),
					Namespace: "",
					Types:     "namespaces",
				}).Invoke(c).Suffix("/").ToA()
			case 1:
				return kubectl.ActionResources(kubectl.ResourceOpts{
					Context:   rootCmd.Flag("context").Value.String(),
					Namespace: c.Parts[0],
					Types:     "pods",
				}).Invoke(c).Suffix(":").ToA()
			default:
				return carapace.ActionValues()
			}
		}).UnlessF(condition.CompletingPath),
	).ToA()
}
