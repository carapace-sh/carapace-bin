package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:     "debug (POD | TYPE[[.VERSION].GROUP]/NAME) [ -- COMMAND [args...] ]",
	Short:   "Create debugging sessions for troubleshooting workloads and nodes",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()

	debugCmd.Flags().Bool("arguments-only", false, "If specified, everything after -- will be passed to the new container as Args instead of Command.")
	debugCmd.Flags().Bool("attach", false, "If true, wait for the container to start running, and then attach as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true.")
	debugCmd.Flags().StringP("container", "c", "", "Container name to use for debug container.")
	debugCmd.Flags().String("copy-to", "", "Create a copy of the target Pod with this name.")
	debugCmd.Flags().String("custom", "", "Path to a JSON or YAML file containing a partial container spec to customize built-in debug profiles.")
	debugCmd.Flags().String("env", "", "Environment variables to set in the container.")
	debugCmd.Flags().StringSliceP("filename", "f", nil, "identifying the resource to debug")
	debugCmd.Flags().String("image", "", "Container image to use for debug container.")
	debugCmd.Flags().String("image-pull-policy", "", "The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server.")
	debugCmd.Flags().Bool("keep-annotations", false, "If true, keep the original pod annotations.(This flag only works when used with '--copy-to')")
	debugCmd.Flags().Bool("keep-init-containers", false, "Run the init containers for the pod. Defaults to true.(This flag only works when used with '--copy-to')")
	debugCmd.Flags().Bool("keep-labels", false, "If true, keep the original pod labels.(This flag only works when used with '--copy-to')")
	debugCmd.Flags().Bool("keep-liveness", false, "If true, keep the original pod liveness probes.(This flag only works when used with '--copy-to')")
	debugCmd.Flags().Bool("keep-readiness", false, "If true, keep the original pod readiness probes.(This flag only works when used with '--copy-to')")
	debugCmd.Flags().Bool("keep-startup", false, "If true, keep the original startup probes.(This flag only works when used with '--copy-to')")
	debugCmd.Flags().String("profile", "", "Options are \"legacy\", \"general\", \"baseline\", \"netadmin\", \"restricted\" or \"sysadmin\".")
	debugCmd.Flags().BoolP("quiet", "q", false, "If true, suppress informational messages.")
	debugCmd.Flags().Bool("replace", false, "When used with '--copy-to', delete the original Pod.")
	debugCmd.Flags().Bool("same-node", false, "When used with '--copy-to', schedule the copy of target Pod on the same node.")
	debugCmd.Flags().String("set-image", "", "When used with '--copy-to', a list of name=image pairs for changing container images, similar to how 'kubectl set image' works.")
	debugCmd.Flags().Bool("share-processes", false, "When used with '--copy-to', enable process namespace sharing in the copy.")
	debugCmd.Flags().BoolP("stdin", "i", false, "Keep stdin open on the container(s) in the pod, even if nothing is attached.")
	debugCmd.Flags().String("target", "", "When using an ephemeral container, target processes in this container name.")
	debugCmd.Flags().BoolP("tty", "t", false, "Allocate a TTY for the debugging container.")
	rootCmd.AddCommand(debugCmd)

	// TODO flag completion
	carapace.Gen(debugCmd).FlagCompletion(carapace.ActionMap{
		"env":     env.ActionNameValues(false),
		"profile": carapace.ActionValues("legacy", "general", "baseline", "netadmin", "restricted", "sysadmin"),
	})
}
