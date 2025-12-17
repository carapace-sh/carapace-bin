package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run NAME --image=image [--env=\"key=value\"] [--port=port] [--dry-run=server|client] [--overrides=inline-json] [--command] -- [COMMAND] [args...]",
	Short:   "Run a particular image on the cluster",
	GroupID: "basic beginner",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	runCmd.Flags().StringSlice("annotations", nil, "Annotations to apply to the pod.")
	runCmd.Flags().Bool("attach", false, "If true, wait for the Pod to start running, and then attach to the Pod as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true. With '--restart=Never' the exit code of the container process is returned.")
	runCmd.Flags().String("cascade", "", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.")
	runCmd.Flags().Bool("command", false, "If true and extra arguments are present, use them as the 'command' field in the container, rather than the 'args' field which is the default.")
	runCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	runCmd.Flags().StringSlice("env", nil, "Environment variables to set in the container.")
	runCmd.Flags().Bool("expose", false, "If true, create a ClusterIP service associated with the pod.  Requires `--port`.")
	runCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	runCmd.Flags().StringSliceP("filename", "f", nil, "to use to replace the resource.")
	runCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.")
	runCmd.Flags().String("grace-period", "", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).")
	runCmd.Flags().String("image", "", "The image for the container to run.")
	runCmd.Flags().String("image-pull-policy", "", "The image pull policy for the container.  If left empty, this value will not be specified by the client and defaulted by the server.")
	runCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	runCmd.Flags().StringP("labels", "l", "", "Comma separated labels to apply to the pod. Will override previous values.")
	runCmd.Flags().Bool("leave-stdin-open", false, "If the pod is started in interactive mode or with stdin, leave stdin open after the first attach completes. By default, stdin will be closed after the first attach completes.")
	runCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	runCmd.Flags().String("override-type", "", "The method used to override the generated object: json, merge, or strategic.")
	runCmd.Flags().String("overrides", "", "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.")
	runCmd.Flags().String("pod-running-timeout", "", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running")
	runCmd.Flags().String("port", "", "The port that this container exposes.")
	runCmd.Flags().Bool("privileged", false, "If true, run the container in privileged mode.")
	runCmd.Flags().BoolP("quiet", "q", false, "If true, suppress prompt messages.")
	runCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	runCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	runCmd.Flags().String("restart", "", "The restart policy for this Pod.  Legal values [Always, OnFailure, Never].")
	runCmd.Flags().Bool("rm", false, "If true, delete the pod after it exits.  Only valid when attaching to the container, e.g. with '--attach' or with '-i/--stdin'.")
	runCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	runCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	runCmd.Flags().BoolP("stdin", "i", false, "Keep stdin open on the container in the pod, even if nothing is attached.")
	runCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	runCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a TTY for the container in the pod.")
	runCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	runCmd.Flag("cascade").NoOptDefVal = " "
	runCmd.Flag("dry-run").NoOptDefVal = " "
	runCmd.MarkFlagRequired("image")
	runCmd.Flag("record").Hidden = true
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":           kubectl.ActionDryRunModes(),
		"env":               env.ActionNameValues(false),
		"image-pull-policy": carapace.ActionValues("Never", "Always", "IfNotPresent"),
		"output":            kubectl.ActionOutputFormats(),
		"restart":           carapace.ActionValues("Always", "OnFailure", "Never"),
		"template":          carapace.ActionFiles(),
	})
}
