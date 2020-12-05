package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a particular image on the cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("allow-missing-template-keys", false, "ignore any errors in templates when a key is missing")
	runCmd.Flags().Bool("attach", false, "attach to the Pod")
	runCmd.Flags().Bool("cascade", false, "cascade the deletion of the resources")
	runCmd.Flags().Bool("command", false, "use extra arguments are present as the 'command' field in the container")
	runCmd.Flags().String("dry-run", "", "run without sending or persisting")
	runCmd.Flags().Bool("expose", false, "service is created for the container(s) which are run")
	runCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership")
	runCmd.Flags().String("filename", "", "to use to replace the resource.")
	runCmd.Flags().Bool("force", false, "immediately remove resources from API and bypass graceful deletion")
	runCmd.Flags().Bool("grace-period", false, "PERIOD    Period of time in seconds given to the resource to terminate gracefully")
	runCmd.Flags().String("hostport", "", "The host port mapping for the container port")
	runCmd.Flags().String("image", "", "The image for the container to run")
	runCmd.Flags().String("image-pull-policy", "", "The image pull policy for the container")
	runCmd.Flags().String("kustomize", "", "Process a kustomization directory")
	runCmd.Flags().String("labels", "", "Comma separated labels to apply to the pod(s)")
	runCmd.Flags().Bool("leave-stdin-open", false, "leave stdin open after the first attach completes")
	runCmd.Flags().String("limits", "", "The resource requirement limits for this container")
	runCmd.Flags().String("output", "", "Output format")
	runCmd.Flags().String("overrides", "", "An inline JSON override for the generated object")
	runCmd.Flags().String("pod-running-timeout", "", "The length of time to wait until at least one pod is running")
	runCmd.Flags().String("port", "", "The port that this container exposes")
	runCmd.Flags().Bool("privileged", false, "run the container in privileged mode")
	runCmd.Flags().Bool("quiet", false, "suppress prompt messages")
	runCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation")
	runCmd.Flags().Bool("recursive", false, "Process the directory used in -f, --filename recursively")
	runCmd.Flags().String("requests", "", "The resource requirement requests for this container")
	runCmd.Flags().String("restart", "", "The restart policy for this Pod")
	runCmd.Flags().Bool("rm", false, "delete resources created in this command for attached containers")
	runCmd.Flags().Bool("save-config", false, "the configuration of current object will be saved in its annotation")
	runCmd.Flags().String("serviceaccount", "", "Service account to set in the pod spec")
	runCmd.Flags().Bool("stdin", false, "Keep stdin open on the container(s) in the pod")
	runCmd.Flags().String("template", "", "Template string or path to template file to use")
	runCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete")
	runCmd.Flags().Bool("tty", false, "Allocated a TTY for each container in the pod")
	runCmd.Flags().Bool("wait", false, "wait for resources to be gone before returning")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":           carapace.ActionValues("none", "server", "client"),
		"image-pull-policy": carapace.ActionValues("Never", "Always", "IfNotPresent"),
		"kustomize":         carapace.ActionDirectories(),
		"output":            carapace.ActionValues("json", "yaml", "name", "go-template", "go-template-file", "template", "templatefile", "jsonpath", "jsonpath-as-json", "jsonpath-file"),
		"restart":           carapace.ActionValues("Always", "OnFailure", "Never"),
		"template":          carapace.ActionFiles(""),
	})
}
