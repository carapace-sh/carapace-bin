package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a Kubernetes pod.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_execCmd).Standalone()

	kube_execCmd.Flags().StringP("container", "c", "", "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen")
	kube_execCmd.Flags().StringP("filename", "f", "", "to use to exec into the resource")
	kube_execCmd.Flags().String("invite", "", "A comma separated list of people to mark as invited for the session.")
	kube_execCmd.Flags().Bool("no-participant-req", false, "Displays a verbose list of required participants in a moderated session.")
	kube_execCmd.Flags().Bool("no-quiet", false, "Only print output from the remote session")
	kube_execCmd.Flags().Bool("no-stdin", false, "Pass stdin to the container")
	kube_execCmd.Flags().Bool("no-tty", false, "Stdin is a TTY")
	kube_execCmd.Flags().Bool("participant-req", false, "Displays a verbose list of required participants in a moderated session.")
	kube_execCmd.Flags().BoolP("quiet", "q", false, "Only print output from the remote session")
	kube_execCmd.Flags().String("reason", "", "The purpose of the session.")
	kube_execCmd.Flags().BoolP("stdin", "s", false, "Pass stdin to the container")
	kube_execCmd.Flags().BoolP("tty", "t", false, "Stdin is a TTY")
	kube_execCmd.Flag("no-participant-req").Hidden = true
	kube_execCmd.Flag("no-quiet").Hidden = true
	kube_execCmd.Flag("no-stdin").Hidden = true
	kube_execCmd.Flag("no-tty").Hidden = true
	kubeCmd.AddCommand(kube_execCmd)
}
