package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec (POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...]",
	Short:   "Execute a command in a container",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().StringP("container", "c", "", "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen")
	execCmd.Flags().StringSliceP("filename", "f", nil, "to use to exec into the resource")
	execCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	execCmd.Flags().BoolP("quiet", "q", false, "Only print output from the remote session")
	execCmd.Flags().BoolP("stdin", "i", false, "Pass stdin to the container")
	execCmd.Flags().BoolP("tty", "t", false, "Stdin is a TTY")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).FlagCompletion(carapace.ActionMap{
		"container": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionMessage("no resource specified")
			} else {
				return kubectl.ActionContainers(kubectl.ContainerOpts{
					Namespace: rootCmd.Flag("namespace").Value.String(),
					Resource:  c.Args[0],
				})
			}
		}),
	})

	carapace.Gen(execCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionApiResourceResources(kubectl.ApiResourceResourcesOpts{
				Namespace: rootCmd.Flag("namespace").Value.String(),
			})
		}),
	)
}
