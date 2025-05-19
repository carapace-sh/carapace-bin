package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var replaceCmd = &cobra.Command{
	Use:     "replace -f FILENAME",
	Short:   "Replace a resource by file name or stdin",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(replaceCmd).Standalone()

	replaceCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	replaceCmd.Flags().String("cascade", "", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.")
	replaceCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	replaceCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	replaceCmd.Flags().StringSliceP("filename", "f", nil, "The files that contain the configurations to replace.")
	replaceCmd.Flags().Bool("force", false, "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.")
	replaceCmd.Flags().String("grace-period", "", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).")
	replaceCmd.Flags().StringP("kustomize", "k", "", "Process a kustomization directory. This flag can't be used together with -f or -R.")
	replaceCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	replaceCmd.Flags().String("raw", "", "Raw URI to PUT to the server.  Uses the transport specified by the kubeconfig file.")
	replaceCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	replaceCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	replaceCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	replaceCmd.Flags().String("subresource", "", "If specified, replace will operate on the subresource of the requested object.")
	replaceCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	replaceCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object")
	replaceCmd.Flags().String("validate", "", "Must be one of: strict (or true), warn, ignore (or false). \"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.")
	replaceCmd.Flags().Bool("wait", false, "If true, wait for resources to be gone before returning. This waits for finalizers.")
	replaceCmd.Flag("cascade").NoOptDefVal = " "
	replaceCmd.Flag("dry-run").NoOptDefVal = " "
	replaceCmd.Flag("validate").NoOptDefVal = " "
	rootCmd.AddCommand(replaceCmd)

	carapace.Gen(replaceCmd).FlagCompletion(carapace.ActionMap{
		"cascade":   carapace.ActionValues("background", "orphan", "foreground"),
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"validate":  kubectl.ActionValidationModes(),
	})
}
